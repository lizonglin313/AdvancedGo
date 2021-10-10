package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string   `json:"user"`
	Threads []string `json:"threads"`
}

// header
// @Desc:
// @Param:	w
// @Param:	r
// @Notice:	WriteHeader 执行完毕后 不允许再写入信息了
func header(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.Header().Set("A", "aaa")
	w.WriteHeader(302)
}

// outputJSON
// @Desc:
// @Param:	w
// @Param:	r
// @Notice:	要首先 使用 Header 将相应的内容 设置为 JSON 格式
func outputJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "lzl",
		Threads: []string{"1", "2", "3"},
	}
	jsonByte, _ := json.Marshal(post)
	w.Write(jsonByte)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/header", header)
	http.HandleFunc("/json", outputJSON)

	server.ListenAndServe()
}
