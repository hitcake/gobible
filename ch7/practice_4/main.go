package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
)

type StringReader struct {
	s   string
	cur int
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	if sr.cur >= len(sr.s) {
		return 0, io.EOF
	}
	n = copy(p, sr.s[sr.cur:])
	sr.cur += n
	return
}

func NewReader(s string) io.Reader {
	return &StringReader{s: s, cur: 0}
}

func main() {

	htmlStr := `<html><body><h1>Hello, Go!</h1><p>This is a test.</p></body></html>`

	// 使用 NewReader 构造 io.Reader
	reader := NewReader(htmlStr)

	// 使用 html.Parse 解析 HTML
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 HTML 节点并打印标签名
	var visit func(*html.Node)
	visit = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Println("Tag:", n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c)
		}
	}
	visit(doc)
}
