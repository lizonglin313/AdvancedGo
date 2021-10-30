package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 设置静态文件加载
	router := gin.Default()
	// relativePath 是 html 模板文件里的前缀 然后替换成 root 表示的相对于项目根的全路径
	router.Static("../static", "gin_demo/serve_files/static")
	// router.StaticFS() // 更换了文件系统（可能是这个意思 用了 http.FileSystem
	// router.StaticFile() // 只是加载指定的某个或者某几个文件
	router.LoadHTMLGlob("./gin_demo/serve_files/temp/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.GET("/local/file", func(c *gin.Context) {
		c.File("gin_demo/serve_files/serve_files.go")
	})

	//var fs http.FileSystem = // ...
	//router.GET("/fs/file", func(c *gin.Context) {
	//	c.FileFromFS("fs/file.go", fs)
	//})

	router.GET("/datafromreader", func(c *gin.Context) {
		response, err := http.Get("http://127.0.0.1:8123")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		defer reader.Close()

		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.jpeg"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

	})
	router.Run(":8123")
}
