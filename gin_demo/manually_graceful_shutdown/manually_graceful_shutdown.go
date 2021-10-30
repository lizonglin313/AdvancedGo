package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		time.Sleep(20 * time.Second)
		c.String(http.StatusOK, "welcome to gin server")
	})

	srv := &http.Server{
		Addr:    ":8123",
		Handler: router,
	}

	// 在 goroutine 中初始化服务 所以它不会阻塞接下来的优雅关闭
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
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
	// 通俗的将，就是暂时把某个请求放过去，然后告诉它，再给它5秒时间去处理完业务
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// shutdown 貌似是 go1.8 之后增加的用于服务器优雅关停的方法
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}














