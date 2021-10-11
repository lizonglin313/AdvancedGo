package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// 使用 encoding/csv 进行 CSV的读写

// Post
// @Description: 文章
type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	// 创建文件
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// 构造数据
	allPosts := []Post{
		Post{Id: 1, Content: "post1", Author: "lzl"},
		Post{Id: 2, Content: "post2", Author: "dkc"},
		Post{Id: 3, Content: "post3", Author: "bbb"},
		Post{Id: 4, Content: "post4", Author: "lzl"},
	}

	// 创建 writer 进行写入
	// 注意 Write 方法一开始是写入 buffer 中的
	// 所以最后必须用 Flush 将数据从 buffer 写入 文件
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	//=============================================//
	// 打开 CSV
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建 reader 进行读取
	reader := csv.NewReader(file)
	// FieldsPerRecord 表示每条记录的预期的字段数
	// > 0 表示预期的字段
	// = 0 表示默认为第一条记录的字段数
	// < 0 表示字段数可变
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// 从 record 中拿
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{
			Id: int(id),
			Content: item[1],
			Author: item[2],
		}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}
