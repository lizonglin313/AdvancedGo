package main

import (
	"fmt"
	"net/http"
)

// 处理器和处理器函数

type HelloHandler struct {}	// 这是一个处理器
type WorldHandler struct {}	// 同样的也是一个处理器

// 为它们定义动作，实现 Handler 中的 ServeHttp 方法
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World!")
}

// 使用处理器函数 处理请求
func helloFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello function!")
}

func main() {
	// 使用处理器
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 直接绑定 处理器
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	// 使用 处理函数
	http.HandleFunc("/hellof", helloFunc)

	server.ListenAndServe()

}
