package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 读写文件

func main() {
	data := []byte("Hello World!\n")

	// 通过 WriteFile 和 ReadFile 对文件进行读写
	// 直接通过 文件名 进行读写
	err := ioutil.WriteFile("data1", data, 0664)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	// 通过 File 结构对文件进行读写
	// 通过 Create - Write 写文件，Open - Read 读文件
	// 都需要拿到文件指针进行操作
	// 比较繁琐，但是某些情况下读取文件更加灵活
	file1, _ := os.Create("data2")
	defer file1.Close()

	byteNum, _ := file1.Write(data)
	fmt.Printf("Write %d bytes to file!\n", byteNum)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	byteNum, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file!\n", byteNum)
	fmt.Println(string(read2))

}
