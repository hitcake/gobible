package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	cancel := make(chan struct{})
	responses := make(chan *http.Response)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go REQUEST_URI")
		os.Exit(1)
	}
	url := os.Args[1]
	go func() {
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			panic(err)
		}
		request.Cancel = cancel
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	ticker := time.Tick(3 * time.Second)
	select {
	case resp := <-responses:
		defer resp.Body.Close()
		result, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(result))
	case <-ticker:
		close(cancel)
		fmt.Printf("%s timeout", url)
	}

}
