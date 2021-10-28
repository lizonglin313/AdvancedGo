package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20

	router.LoadHTMLGlob("./gin_demo/multiple_files/public/*")

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err %s", err.Error()))
			return
		}

		files := form.File["files"]
		for _, file := range files {
			filename := "gin_demo/multiple_files/" + filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	})

	router.Run(":8123")
}
