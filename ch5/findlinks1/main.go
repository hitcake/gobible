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
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			continue
		}
		for _, link := range visit(nil, doc) {
			fmt.Println(link)
		}
	}

}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
