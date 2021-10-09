package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func processFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	// 通过 form 表单中的 name 字段去取
	fileHeader := r.MultipartForm.File["uploaded"][0]
	fmt.Println(fileHeader)
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

// formFile
// @Desc: 	和 FormValue 一样 使用 FormFile
//			自动调用 ParseMultipartForm 获得第一个值（文件）
// @Param:	w
// @Param:	r
// @Notice:
func formFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintf(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/processFile", processFile)

	server.ListenAndServe()
}
