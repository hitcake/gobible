package main

import (
	"bytes"
	"fmt"
	"gobible/ch5/fetch"
	"golang.org/x/net/html"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		data, err := fetch.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		doc, err := html.Parse(bytes.NewReader(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing %s as HTML: %v\n", url, err)
			continue
		}

		fmt.Println(extractText(doc))
	}
}

func extractText(n *html.Node) (text []string) {
	if n.Type == html.TextNode {
		text = append(text, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}
		text = append(text, extractText(c)...)
	}
	return
}
