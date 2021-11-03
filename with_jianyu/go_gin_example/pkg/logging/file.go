package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 定义打logging的基本属性
var (
	LogSavePath = "with_jianyu/go_gin_example/runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// getLogFileFullPath
// @Desc: 	获取日志全路径
// @Return:	string
// @Notice:
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

// openLogFile
// @Desc: 	通过文件路径返回文件句柄
// @Param:	filePath
// @Return:	*os.File
// @Notice:
func openLogFile(filePath string) *os.File {
	// 返回文件信息结构描述文件
	// 如果出现错误，会返回*PathError
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile: %v", err)
	}
	return handle
}

func mkDir() {
	// os.Getwd：返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/runtime/logs/", os.ModePerm)
	if err != nil {
		panic(err)
	}
}
