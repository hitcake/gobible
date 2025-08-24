package main

import (
	"fmt"
	"gobible/ch5/links"
	"os"
	"sync"
	"time"
)

var tokens = make(chan struct{}, 20)

const MAX_DEPTH = 3

type DepthUrl struct {
	url   string
	depth int
}

func crawl(du DepthUrl) ([]DepthUrl, error) {
	if du.depth > MAX_DEPTH {
		return nil, fmt.Errorf("depth too high")
	}
	fmt.Printf("url %s, depth %d\n", du.url, du.depth)
	tokens <- struct{}{}
	list, err := links.Extract(du.url)
	<-tokens
	if err != nil {
		return nil, err
	}
	result := make([]DepthUrl, 0, len(list))
	for _, v := range list {
		result = append(result, DepthUrl{url: v, depth: du.depth + 1})
	}
	return result, nil
}

func main() {
	worklist := make(chan []DepthUrl)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		list := make([]DepthUrl, 0, len(os.Args[1:]))
		for _, url := range os.Args[1:] {
			list = append(list, DepthUrl{url: url, depth: 1})
		}
		worklist <- list
	}()
	go func() {
		time.Sleep(10 * time.Second)
		wg.Wait()
		close(worklist)
	}()
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				wg.Add(1)
				go func(du DepthUrl) {
					defer wg.Done()
					list, err := crawl(du)
					if err == nil {
						worklist <- list
					}
				}(link)
			}
		}
	}

}
