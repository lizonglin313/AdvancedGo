package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 在中间件中使用 gouroutine 需要使用 context 的副本

func main() {

	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()	// 把 context 做拷贝
		go func() {
			time.Sleep(5 * time.Second)

			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
		log.Println("open go func in async then go to sync")
		c.Redirect(302, "/long_sync")
	})

	r.GET("/long_sync", func(c *gin.Context) {
		log.Println("in sync ")
		time.Sleep(5 * time.Second)

		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run(":8123")
}
