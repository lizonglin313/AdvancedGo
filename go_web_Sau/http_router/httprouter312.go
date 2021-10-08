/**
 @author: zonglin
 @date: 2021/10/8
 @note: 使用httprouter
**/
package http_router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func doHttpRouterDemo()  {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,		// 不再使用默认的多路复用器，而是使用 http router 提供的
	}
	server.ListenAndServe()
}
