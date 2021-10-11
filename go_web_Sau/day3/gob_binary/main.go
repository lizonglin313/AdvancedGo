package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

// 使用 gob 读写 二进制

// Post
// @Description: 文章
type Post struct {
	Id      int
	Content string
	Author  string
}

// store
// @Desc: 	写文件
// @Param:	data
// @Param:	filename
// @Notice:
func store(data interface{}, filename string) {
	// 创建 缓冲区
	buffer := new(bytes.Buffer)

	// 创建 编码器 向缓冲区写
	encoder := gob.NewEncoder(buffer)

	// 编码数据 写道 缓冲区 data -> buffer
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	// 将 缓冲区数据 写到 文件
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// load
// @Desc: 	从 文件中 载入数据
// @Param:	data
// @Param:	filename
// @Notice:
func load(data interface{}, filename string) {
	// 从 文件中 读取原始数据
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// 创建缓冲区 将原始数据写到缓冲区 raw -> buffer
	buffer := bytes.NewBuffer(raw)

	// 为 缓冲区 创建 解码器
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "post1", Author: "lzl"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
