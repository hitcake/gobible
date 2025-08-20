package main

import (
	"bytes"
	"fmt"
	"gobible/ch5/fetch"
	"golang.org/x/net/html"
	"os"
)

var nodelink = map[string]string{"a": "href", "script": "src", "style": "src"}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findlinks(url string) ([]string, error) {
	data, err := fetch.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch: %s %v\n", url, err)
	}
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		attr := nodelink[n.Data]
		if attr != "" {
			for _, a := range n.Attr {
				if a.Key == attr {
					links = append(links, a.Val)
				}
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
