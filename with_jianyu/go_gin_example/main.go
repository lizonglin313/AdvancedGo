package main

import (
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"AdvancedGo/with_jianyu/go_gin_example/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	s.ListenAndServe()
}
