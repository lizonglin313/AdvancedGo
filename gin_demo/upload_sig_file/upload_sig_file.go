package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

// 1000

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 一般需要先加载如 js css 等静态文件
	// 加载全局模板文件， 路径是相对于 根目录的相对路径
	router.LoadHTMLGlob("./gin_demo/upload_sig_file/public/*")
	//router.Static("/gin_demo/upload_sig_file/public", "./gin_demo/upload_sig_file/public")

	router.GET("/upload", func(c *gin.Context) {
		// 返回 html 模板
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		// 这样就会存到 gin_demo 下面
		filename := "gin_demo/" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	})
	router.Run(":8123")
}
