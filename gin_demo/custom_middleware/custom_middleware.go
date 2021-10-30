package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")
		log.Print("set example in middlerware")
		// 以上动作在实际请求前
		// 当前中间件中调用c.Next()时会中断当前中间件中后续的逻辑
		// 转而执行后续的中间件和handlers
		// 等他们全部执行完以后再回来执行当前中间件的后续代码
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		log.Print("in /test handler")
		// 使用断言明确数据类型
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8123")
}
