package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
