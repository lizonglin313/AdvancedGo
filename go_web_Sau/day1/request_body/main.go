package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.Header)
	len := r.ContentLength
	body := make([]byte, len)
	// 将 body 写入
	r.Body.Read(body)
	fmt.Printf(string(body))
	fmt.Fprintf(w, string(body))

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
