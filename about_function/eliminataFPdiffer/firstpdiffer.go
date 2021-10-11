package main

import (
	"fmt"
	"os"
)

// 消除首参差异

// 1.不依赖具体对象
func doFile() {
	var (
		CloseFile = (*os.File).Close	// 变量就是一个方法
		ReadFile  = (*os.File).Read
	)

	f, _ := os.OpenFile("filename.txt",os.O_RDWR|os.O_CREATE, 0755)
	ReadFile(f, []byte("data"))
	CloseFile(f)	// 如果直接用 f 调用 -- f.Close()
}

// 2.通过闭包去除参数 f 消除首参差异
func doFile1() {
	f, _ := os.OpenFile("filename.txt", os.O_RDWR|os.O_CREATE, 0755)
	// 通常的，匿名函数用来直接使用函数外的变量，利用的就是这一个特性
	var (
		Close = func() error {
		return (*os.File).Close(f)
	}
		Read = func(data []byte) (int, error) {
			return (*os.File).Read(f, data)
		}
	)

	// 不再依赖 f，只要是个 *File 类型的都可以放
	Read([]byte("data"))
	Close()

}

// 3.用方法值简化， 绑定
func doFile2() {
	f, _ := os.OpenFile("filename.txt", os.O_RDWR|os.O_CREATE, 0755)

	var (
		Close = f.Close
		Read = f.Read
	)

	// 处理
	Read([]byte("data"))
	Close()

}

func main() {

	fmt.Println()
}
