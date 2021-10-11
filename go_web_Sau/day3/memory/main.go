package main

import "fmt"

// 使用内存进行存储

// Post
// @Description: 定义文章结构体
type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post           // Id -> Post 的映射
var PostsByAuthor map[string][]*Post // string -> []*Post 的映射

// store
// @Desc: 	存储文章
// @Param:	post
// @Notice:	同时向不同映射中存储，注意 PostsByAuthor 中 value 是一个 切片
func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	// 实例化定义
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "post1", Author: "lzl"}
	post2 := Post{Id: 2, Content: "post2", Author: "dkc"}
	post3 := Post{Id: 3, Content: "post3", Author: "bbb"}
	post4 := Post{Id: 4, Content: "post4", Author: "lzl"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["lzl"]{
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["dkc"]{
		fmt.Println(post)
	}

}
