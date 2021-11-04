package main

import (
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"AdvancedGo/with_jianyu/go_gin_example/routers"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// 等待结束的中断信号 (5s之后
	quit := make(chan os.Signal)
	// kill 信号默认发送给 syscall.SIGTERM
	// kill -2 是 syscall.SIGINT
	// kill -9 是 syscall.SIGKILL 但是不能被捕获
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)	// 把信号传给 channel
	<-quit
	log.Println("Shutting down server")

	// 上下文 context 被用来通知5秒后结束服务
	// 请求目前还是正常处理的
	// 通俗的讲，就是暂时把某个请求放过去，然后告诉它，再给它5秒时间去处理完业务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
