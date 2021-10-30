package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// group v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			name := c.PostForm("name")
			c.String(http.StatusOK, "hello %s ,this is v1's login", name)
		})
		v1.GET("/creat/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "creat %s using v1 successful!", name)
		})
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			name := c.PostForm("name")
			c.String(http.StatusOK, "hello %s ,this is v2's login", name)
		})
		v2.GET("/creat/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "creat %s using v2 successful!", name)
		})
	}
	router.Run(":8123")
}
