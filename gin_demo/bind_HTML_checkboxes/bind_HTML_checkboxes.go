package main

import "github.com/gin-gonic/gin"

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("gin_demo/bind_HTML_checkboxes/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "")
	})

	r.POST("/", formHandler)

	r.Run(":8123")
}
