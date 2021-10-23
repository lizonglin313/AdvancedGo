package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//var depth int
//
//func startElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		// 使用 * 填充一些数量的空格 也就和深度有关
//		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
//		depth++
//	}
//}
//
//func endElement(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth--
//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
//	}
//}

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

// 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素，因为这些元素对浏览者是不可见的
func getText(s string) {
	// s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func main() {
	url := "https://www.liwenzhou.com/"
	context := Fetch(url)

	doc, err := html.Parse(strings.NewReader(context))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	depth := 0
	startFunc := func(n *html.Node) {
		if n.Type == html.ElementNode {
			// 使用 * 填充一些数量的空格 也就和深度有关
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}


	endFunc := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, startFunc, endFunc)
}
