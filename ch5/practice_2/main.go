package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	count := make(map[string]int)
	visit(doc, count)
	for k, v := range count {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func visit(n *html.Node, count map[string]int) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, count)
	}
}
