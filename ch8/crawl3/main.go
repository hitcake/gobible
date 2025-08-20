package main

import (
	"fmt"
	"gobible/ch5/links"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseedLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseedLinks {
				foundLink := crawl(link)
				go func() { worklist <- foundLink }()
			}
		}()
	}
	seen := make(map[string]bool)

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseedLinks <- link
			}
		}
	}
}
