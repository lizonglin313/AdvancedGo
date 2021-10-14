package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"time"
	// "runtime"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	// building up the source tile database
	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("go_web_Sau/day5/concurrent_mosaic/upload.html")
	t.Execute(w, nil)
}

//func mosaic(w http.ResponseWriter, r *http.Request) {
//	t0 := time.Now()
//	// get the content from the POSTed form
//	r.ParseMultipartForm(10485760) // 设置文件最大 10 MB
//	file, _, _ := r.FormFile("image")
//	defer file.Close()
//	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
//	// decode and get original image
//	original, _, _ := image.Decode(file)
//	bounds := original.Bounds()
//	db := cloneTilesDB()
//
//	// 切分处理马赛克
//	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
//	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
//	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
//	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)
//
//	// 合并
//	c := conbine(bounds, c1, c2, c3, c4)
//
//	buf1 := new(bytes.Buffer)
//	jpeg.Encode(buf1, original,nil)
//	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())
//
//
//	t1 := time.Now()
//	images := map[string]string{
//		"original": originalStr,
//		"mosaic":   <-c,
//		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
//	}
//	t, _ := template.ParseFiles("go_web_Sau/day5/concurrent_mosaic/results.html")
//	t.Execute(w, images)
//
//}

func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	// get the content from the POSTed form
	r.ParseMultipartForm(10485760) // max body in memory is 10MB
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	//
	//   // decode and get original image
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := cloneTilesDB()

	// fan-out
	c1 := cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)
	// fan-in
	c := combine(bounds, c1, c2, c3, c4)
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   <-c,
		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}

	t, _ := template.ParseFiles("go_web_Sau/day5/concurrent_mosaic/results.html")
	t.Execute(w, images)
}
