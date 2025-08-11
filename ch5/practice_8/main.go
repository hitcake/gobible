package main

import (
	"GoBible/ch5/fetch"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := forEachNode(c, id, pre, post); found != nil {
			return found
		}
	}
	if post != nil {
		if !post(n, id) {
			return n
		}
	}
	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}

func endElement(n *html.Node, id string) bool {
	return true
}
func ElementByID(n *html.Node, id string) *html.Node {
	return forEachNode(n, id, startElement, endElement)
}
func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <url> <id>\n", os.Args[0])
		return
	}
	data, err := fetch.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	if p := ElementByID(doc, os.Args[2]); p != nil {
		fmt.Printf("Found %s:", p.Data)
		for _, a := range p.Attr {
			fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
		}
	} else {
		fmt.Println("Not found.")
	}

}
