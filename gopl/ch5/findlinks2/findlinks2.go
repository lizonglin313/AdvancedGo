package main

import (
	"AdvancedGo/log_demo"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findlinks(url string) ([]string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log_demo.Error("can not get url:", url)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log_demo.Error("status code is:", resp.StatusCode)
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log_demo.Error("parse body error:", err)
		return nil, err
	}
	return visit(nil, doc), nil
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

func main() {
	strings, err := findlinks("https://www.liwenzhou.com/")
	if err != nil {
		log_demo.Error("error")
	}
	for _, v := range strings {
		fmt.Println(v)
	}
}
