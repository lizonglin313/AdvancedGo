package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sync"
)

// 与 单任务版不同 并发版 需要将图像进行分块 处理
// 因此增加两个进行 分块 和 合并 的函数

type DB struct {
	// 瓷砖图片用完就删除，并发过程中互斥访问 nearest 加锁
	mutex *sync.Mutex
	store map[string][3]float64
}

// averageColor
// @Desc: 	将给定图片每个像素中的红绿蓝3种颜色相加
//			将这些颜色除以像素总和，得到平均颜色值
// @Param:	img
// @Return:	[3]float64
// @Notice:
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	// 遍历所有像素点
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			// 计算图片的平均颜色
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	// 求平均 并返回
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

// resize
// @Desc: 	将图片缩放至指定宽度
// @Param:	in
// @Param:	newWidth
// @Return:	image.NRGBA
// @Notice:
func resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	width := bounds.Dx()
	ratio := width / newWidth
	// 存疑 懂了... Min.X 和 Min.Y 一样 都是 0 所以作者这里用了 两个 Min.X
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio,
		bounds.Max.X/ratio, bounds.Max.Y/ratio))

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			// >> 8 右移 8 位 也就是除以 2^8
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8),
				uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

// tilesDB
// @Desc: 	对文件进行批处理，计算每个瓷砖的平均颜色
// @Return:	map[string][3]float64
// @Notice: 这计算的是瓷砖的平均颜色，因为需要依据瓷砖的平均颜色对图片进行填充
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db...")
	db := make(map[string][3]float64)
	files, err := ioutil.ReadDir("go_web_Sau/day5/concurrent_mosaic/tiles")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		//name := "go_web_Sau/day5/concurrent_mosaic/tiles/" + f.Name()
		//fmt.Println(f.Name())
		name := filepath.Join("go_web_Sau/day5/concurrent_mosaic/tiles/", f.Name())
		//fmt.Println(name)
		file, err := os.Open(name)
		if err == nil {
			img, _, err := image.Decode(file)
			if err == nil {
				db[name] = averageColor(img)
			} else {
				fmt.Println("error in populating TILEDB:", err, name)
			}
		} else {
			fmt.Println("cannot open file", name, err)
		}
		file.Close()
	}
	fmt.Println("Finished populating tiles db.")
	// 返回一个  filename -> (r, g, b) 的映射
	return db
}

// nearest
// @Desc: 	寻找与目标图片的平均颜色最接近的瓷砖图片
// @Param:	target	目标图片
// @Param:	db		瓷砖
// @Return:	string
// @Notice: 有 4 个 goroutine 在 并发的 计算 nearest 的瓷砖
//			如果有两个计算的相同 会被 使用 删除两次
func (db *DB) nearest(target [3]float64) string {
	var filename string
	db.mutex.Lock() // 加锁

	smallest := 1000000.0

	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(db.store, filename) // 有可能两个goroutine发生同时删除的情况，所以要加锁
	db.mutex.Unlock()          // 删除了就可以解锁了
	return filename
}

func sq(n float64) float64 {
	return n * n
}

// distance
// @Desc: 	计算两个三元组的欧氏距离
// @Param:	p1
// @Param:	p2
// @Return:	float64
// @Notice:
func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

var TILESDB map[string][3]float64

// cloneTilesDB
// @Desc: 	复制马赛克瓷砖平均颜色的副本 不用再一次次计算
// @Return:	map[string][3]float64
// @Notice:
func cloneTilesDB() DB {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}

	tiles := DB{
		mutex: &sync.Mutex{},
		store: db,
	}
	return tiles
}

// cut
// @Desc: 	将图片进行分割，按分割后的块进行mosaic
// @Param:	original
// @Param:	db
// @Param:	tileSize
// @Param:	x1
// @Param:	y1
// @Param:	x2
// @Param:	y2
// @Return:	<-chan
// @Notice:	返回一个管道 这个管道只能向外送 Image
func cut(original image.Image, db *DB, tileSize,
	x1, y1, x2, y2 int) <-chan image.Image {

	// 虽然 c 定义的是一个双向管道 但是作为 <-chan 的管道使用
	c := make(chan image.Image)
	sp := image.Point{0, 0}
	// 创建goroutine进行分割
	// 这样4块分割就可以同时进行，不用等一个进行完再进行另一个了
	go func() {
		newImage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				// 按照 tileSize 划分粒度
				// 在每个粒度中 拿到原始图像color
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}

				// 从 瓷砖中 找出最相似的
				nearest := db.nearest(color)

				file, err := os.Open(nearest)

				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						t := resize(img, tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
						draw.Draw(newImage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("error in decoding nearest:", err, nearest)
					}
				} else {
					// 如果粒度过小 则会造成数量不够 db[name] 为空 就会报错！
					fmt.Println("error opening file when creating mosaic:", err)
				}
				file.Close()
			}
		}
		c <- newImage.SubImage(newImage.Rect)
	}()
	return c
}

// combine
// @Desc: 	将 4 块 mosaic 合并
// @Param:	r
// @Param:	c1
// @Param:	c2
// @Param:	c3
// @Param:	c4
// @Return:	<-chan
// @Notice:
func combine(r image.Rectangle,
	c1, c2, c3, c4 <-chan image.Image) <-chan string {

	c := make(chan string)

	go func() {
		// 使用 waitgroup 是因为：
		// 开了 goroutine 后 goroutine 的生命周期 就不受这个函数的影响了
		// 如果 goroutine 还没有执行完这个函数就结束了
		// 这个函数。自然拿不到 goroutine 的结果，也就无法完成任务
		var wg sync.WaitGroup
		newImage := image.NewNRGBA(r)
		// 复制块
		copy := func(dst draw.Image, r image.Rectangle,
			src image.Image, sp image.Point, index string) {
			draw.Draw(dst, r, src, sp, draw.Src)

			// ===
			//f, err := os.Create("demo" + index + ".jpeg")
			//if err != nil {
			//	panic(err)
			//}
			//err = jpeg.Encode(f, newImage, nil)
			//if err != nil {
			//	panic(err)
			//}
			// ===

			// 记得释放一个 mutex
			wg.Done() // 匿名函数 使用外面的 变量
		}

		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool

		// 使用循环做，保证 4 块都做完
		for {
			select {
			case s1, ok1 = <-c1:
				go copy(newImage, s1.Bounds(), s1,
					image.Point{r.Min.X, r.Min.Y}, "1")
			case s2, ok2 = <-c2:
				go copy(newImage, s2.Bounds(), s2,
					image.Point{r.Max.X / 2, r.Min.Y}, "2")
			case s3, ok3 = <-c3:
				go copy(newImage, s3.Bounds(), s3,
					image.Point{r.Min.X, r.Max.Y / 2}, "3")
			case s4, ok4 = <-c4:
				go copy(newImage, s4.Bounds(), s4,
					image.Point{r.Max.X / 2, r.Max.Y / 2}, "4")
			}
			// 结束条件
			if ok1 && ok2 && ok3 && ok4 {
				break
			}
		}
		// 先阻塞 保证goroutine运行完
		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, newImage, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()
	return c
}
