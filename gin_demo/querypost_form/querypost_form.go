package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		msg := c.DefaultPostForm("msg", "message")

		fmt.Printf("id: %s; pageL %s; name: %s; message: %s\n", id, page, name, msg)
	})
	router.Run(":8123")
}
