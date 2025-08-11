package main

import (
	"GoBible/ch5/fetch"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
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

var depth int
var hasChildren bool

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf(" %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild != nil {
			fmt.Printf(">\n")
		} else {
			fmt.Printf("/>\n")
		}

		depth++

	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
	if n.Type == html.TextNode {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}

	}
}

func main() {
	for _, url := range os.Args[1:] {
		data, err := fetch.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		doc, err := html.Parse(bytes.NewReader(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline: %v\n", err)
			continue
		}
		fmt.Printf("---------- %s ----------\n", url)
		forEachNode(doc, startElement, endElement)
		fmt.Printf("---------------------------------\n")
	}
}
