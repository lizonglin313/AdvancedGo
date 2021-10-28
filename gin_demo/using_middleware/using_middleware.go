package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 一个中间件
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "using mybenchlogger middleware!")
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "using auth middleware!")
	}
}

func main() {
	// gin.Default 是默认的使用了 Logger 和 Recovery 这两个中间件的
	// router := gin.Default()

	r := gin.New() // 默认的不适用任何中间件

	// 全局中间件
	// 默认写到 os.Stdout 中
	r.Use(gin.Logger())
	// 从 panic 中恢复
	r.Use(gin.Recovery())

	// 特定路由的中间件
	r.GET("/benchmark", MyBenchLogger(), func(c *gin.Context) {
		c.String(http.StatusOK, "第二个参数是中间件， 过了中间件之后才是处理路由的Handler")
	})

	// 一个验证授权Group路由组的中间件
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "using authorized middleware in a router group")
		})

		// 再加一个组
		testing := authorized.Group("testing")
		// 实际要访问的路由是：
		// http://127.0.0.1:8123/testing/testing
		testing.GET("/testing", func(c *gin.Context) {
			c.String(http.StatusOK, "hello testing!")
		})
	}
	r.Run(":8123")
}
