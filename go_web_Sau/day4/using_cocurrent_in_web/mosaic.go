package main
//package using_cocurrent_in_web

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"os"
)

// averageColor
// @Desc: 	将给定图片每个像素中的红绿蓝3种颜色相加
//			将这些颜色除以像素总和，得到平均颜色值
// @Param:	img
// @Return:	[3]float64
// @Notice:
func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			// 计算图片的平均颜色
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
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
	ratio := bounds.Dx() / newWidth
	// 存疑 懂了... Min.X 和 Min.Y 一样 都是 0 所以作者这里用了 两个 Min.X
	out := image.NewNRGBA(image.Rect(bounds.Min.X/ratio, bounds.Min.X/ratio,
		bounds.Max.X/ratio, bounds.Max.Y/ratio))

	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Max.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{uint8(r >> 8), uint8(g >> 8),
				uint8(b >> 8), uint8(a >> 8)})
		}
	}
	return *out
}

// tilesDB
// @Desc: 	对文件进行批处理，计算每个图片的平均颜色
// @Return:	map[string][3]float64
// @Notice:
func tilesDB() map[string][3]float64 {
	fmt.Println("Start populating tiles db...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")

	for _, f := range files {
		name := "tiles/" + f.Name()
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
	return db
}

// nearest
// @Desc: 	寻找与目标图片的平均颜色最接近的瓷砖图片
// @Param:	target	目标图片
// @Param:	db		瓷砖
// @Return:	string
// @Notice:
func nearest(target [3]float64, db *map[string][3]float64) string {
	var filename string
	smallest := 1000000.0

	for k, v := range *db {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(*db, filename)
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
// @Desc: 	复制马赛克瓷砖副本 加快速度
// @Return:	map[string][3]float64
// @Notice:
func cloneTilesDB() map[string][3]float64 {
	db := make(map[string][3]float64)
	for k, v := range TILESDB {
		db[k] = v
	}
	return db
}
