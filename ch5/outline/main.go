package main

import (
	"GoBible/ch5/fetch"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
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
		outline(nil, doc)
	}

}
