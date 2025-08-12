package main

import (
	"GoBible/ch5/fetch"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func ElemantsByTagName(n *html.Node, name ...string) []*html.Node {
	var ret []*html.Node
	if n.Type == html.ElementNode {
		for _, c := range name {
			if n.Data == c {
				ret = append(ret, n)
			}

		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, ElemantsByTagName(c, name...)...)
	}
	return ret
}

func main() {
	for _, url := range os.Args[1:] {
		data, err := fetch.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err:%v\n", err)
			continue
		}
		doc, err := html.Parse(bytes.NewReader(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse %s as HTML err:%v\n", url, err)
			continue
		}
		nodes := ElemantsByTagName(doc, "a")
		fmt.Printf("url: %s, a nodes %d\n", url, len(nodes))
		nodes = ElemantsByTagName(doc, "h1", "img")
		fmt.Printf("url: %s, a and img nodes %d\n", url, len(nodes))
	}

}
