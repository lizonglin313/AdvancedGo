package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)	// 对 n 节点 尝试获取该结点下面的 url
	}

	// 对孩子结点 深度搜索 出来后 搜索兄弟节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	// 因为搜索到结点 把相关的 url 插入就可以了 所以之后不再需要进行操作
	if post != nil {
		post(n)
	}
}

func Extract(url string) ([]string, error) {
	// 处理 URL 拿到请求页面的 content
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// 将 resp 的主体进行 html 解析
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// 进行遍历
	var links []string
	// 定义一个匿名函数 需要一个参数 n
	// 如果节点是 元素 并且是 a 标签
	// 遍历它的属性 然后把 url 加到 links 里面
	// 使用 闭包 就是为了让他带着 links 走
	// 这样 不管到页面的哪一个节点 links 都能保存该页面所有的数据
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	// 只需要在遍历结点前操作节点 所以 post 部分传入 nil
	forEachNode(doc, visitNode, nil)
	return links, nil
}
