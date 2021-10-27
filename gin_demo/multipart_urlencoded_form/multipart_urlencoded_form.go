package main

import "github.com/gin-gonic/gin"

// main
// @Desc: 	拿 form 表单中的字段
// @Notice:
func main() {

	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		msg := c.PostForm("msg")
		nick := c.DefaultPostForm("nick", "defaultFormValue")

		c.JSON(200, gin.H{
			"status": "posted",
			"message": msg,
			"nick": nick,
		})
	})
	router.Run(":8123")
}
