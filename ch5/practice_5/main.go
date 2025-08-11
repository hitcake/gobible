package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}

	}
	if n.Type == html.TextNode {
		words += strings.Count(n.Data, " ")
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {

		childWords, childImages := countWordsAndImages(c)
		words += childWords
		images += childImages
	}
	return
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("url: %s words %d images %d", url, words, images)
	}

}
