package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// main
// @Desc: 	路径中的参数
// @Notice:
func main() {

	router := gin.Default()

	// 只会匹配 /user/xxx 严格的， 匹配 /user  会失败
	// /user/ 会成功 只是 name = 空
	router.GET("/user/:name", func(c *gin.Context) {
		// 从 url 获取参数
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 会匹配 /user/xxx/xxx 和 /user/xxx/
	// 如果 action 字段没有 则默认匹配 /user/xxx/
	// http://127.0.0.1:8123/user/lzl/sda
	// lzl is /sda
	// http://127.0.0.1:8123/user/lzl/
	// lzl is /
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		msg := name + " is " + action
		c.String(http.StatusOK, msg)
	})

	// 每个上下文 会保存 匹配的请求
	router.POST("/user/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/user/:name/*action"
		c.String(http.StatusOK, "%t", b)
	})

	// 即使我们前面定义了 /user/:name 但是 /user/groups 还是会被正确的解析
	// groups 不会被识别成 :name 字段
	// 因为路由解析在参数解析之前
	// /groups/(会被定向到groups)  和  /groups 都可以
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "the available groups are [...]")
	})

	router.Run(":8123")
}
