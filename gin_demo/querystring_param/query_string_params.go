package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// main
// @Desc: 	获取 URL 中 ? 后面的字符
// @Notice:
func main() {

	router := gin.Default()

	// http://127.0.0.1:8123/welcome?firstname=li&lastname=zonglin
	// hello li zonglin
	// http://127.0.0.1:8123/welcome/?firstname=li&lastname=zonglin
	// 也可以但是会被重定向

	// http://127.0.0.1:8123/welcome?lastname=zonglin
	// hello defaultname zonglin
	router.GET("/welcome", func(c *gin.Context) {
		// 如果没有该字段就用 default
		firstname := c.DefaultQuery("firstname", "defaultname")
		lastname := c.Query("lastname")
		// c.Request.URL.Query().Get("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})
	router.Run(":8123")

}
