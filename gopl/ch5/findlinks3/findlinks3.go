package main

import (
	"AdvancedGo/gopl/ch5/links"
	"fmt"
	"log"
)

// 广度优先遍历树

func breadFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url) // 拿到该url页面中的所有url
	if err != nil {
		log.Print(err)
	}
	return list
}

// 按照广度优先的顺序爬取页面的url链接
func main() {
	url := []string{
		"https://www.liwenzhou.com",
	}
	breadFirst(crawl, url)
}
