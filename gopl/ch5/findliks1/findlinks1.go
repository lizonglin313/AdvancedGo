package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func Fetch(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	return string(b)

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 深度优先遍历
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 使用递归形式
func visitUsingRecursive(links []string, n *html.Node) []string {

	if n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 深度优先遍历
	links = visitUsingRecursive(links, n.FirstChild)
	links = visitUsingRecursive(links, n.NextSibling)

	return links
}

// 拓展 查看 a script img link 的元素链接
var elem = map[string]int{
	"a":      1,
	"img":    1,
	"script": 1,
	"link":   1,
}

func visitAllUsingRecursive(links []string, n *html.Node) []string {

	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if _, ok := elem[n.Data]; ok {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
	}

	// 深度优先遍历
	links = visitAllUsingRecursive(links, n.FirstChild)
	links = visitAllUsingRecursive(links, n.NextSibling)

	return links
}

func outlinetext(texts []string, n *html.Node) []string {
	if n == nil || n.Data == "script" || n.Data == "style" {
		return texts
	}

	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	texts = outlinetext(texts, n.FirstChild)
	texts = outlinetext(texts, n.NextSibling)

	return texts
}

func main() {
	url := os.Args[1]
	context := Fetch(url)
	//fmt.Println(context)
	doc, err := html.Parse(strings.NewReader(context))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	//for _, link := range visitUsingRecursive(nil, doc) {
	//	fmt.Println(link)
	//}

	// 查看多种节点的 href
	for _, link := range visitAllUsingRecursive(nil, doc) {
		fmt.Println(link)
	}

	// 查看 textnode
	//for _, text := range outlinetext(nil, doc) {
	//	fmt.Println(text)
	//}
}
