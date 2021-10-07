package chaining_func

import (
"fmt"
"net/http"
"reflect"
"runtime"
)


func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

//
// @Name:	log
// @Desc: 	记录一个handler function的日志
// @Param:	h
// @Return:	http.HandlerFunc
// @Notice: 使用 runtime.FuncForPC 获取正在调用系统日志，也就是该函数的一些信息
//
func log(h http.HandlerFunc) http.HandlerFunc {
	// 使用匿名函数进行返回
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Hanlder function called - " + name)
		h(w, r)
	}
}

// ChainingTwoFunc
// @Desc: 	串联两个处理器函数
// @Notice:
func chainingTwoFunc() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// log 返回一个 handlerfunc 所以它满足类型 handler
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}

