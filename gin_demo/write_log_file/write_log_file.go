package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

const FILEPATH = "gin_demo/write_log_file/log/"

func main() {
	// 设置终端的日志颜色
	// gin.DisableConsoleColor() 这样就没有
	gin.ForceConsoleColor()

	// 把日志写到文件中
	filename := FILEPATH + "gin.log"

	// 创建日志文件
	// f, _ := os.Create(filename)
	// 以追加形式写入
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("create log file error!")
	}

	// 设置Writer
	// gin.DefaultWriter = io.MultiWriter(f)	// 输出目标为文件 f
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // 同时向 文件f 和 终端写入

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8123")
}
