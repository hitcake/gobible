package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	// 设置超时
	client := &http.Client{
		Timeout: time.Second * 10, // 设置超时时间为10秒
	}
	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//ioutil.Discard被弃用
	//nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	nbytes, err := io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

/*
output
1.30    4154 http://gopl.io
1.49   62967 http://go.dev
Get "http://golang.org": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
10.00s elapsed
*/
