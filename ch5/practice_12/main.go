package main

import (
	"bytes"
	"fmt"
	"gobible/ch5/fetch"
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
		depth := 0
		forEachNode(doc, func(n *html.Node) {
			if n.Type == html.ElementNode {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
				depth++
			}
		}, func(n *html.Node) {
			if n.Type == html.ElementNode {
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			}
		})
		fmt.Printf("---------------------------------\n")
	}
}
