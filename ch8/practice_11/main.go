package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	resp, err := mirroredQuery(3 * time.Second)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
func mirroredQuery(duration time.Duration) (string, error) {
	responses := make(chan string, 3)
	tick := time.Tick(duration)
	cancel := make(chan struct{})
	go func() {
		resp, err := request("https://www.baidu.com", cancel)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	go func() {
		resp, err := request("https://europe.gopl.io", cancel)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	go func() {
		resp, err := request("https://americas.gopl.io", cancel)
		if err != nil {
			fmt.Println(err)
		} else {
			responses <- resp
		}
	}()
	select {
	case resp := <-responses:
		close(cancel)
		return resp, nil
	case <-tick:
		close(cancel)
		return "", fmt.Errorf("Timed out")
	}
}

func request(hostname string, cancel <-chan struct{}) (string, error) {
	req, err := http.NewRequest(http.MethodGet, hostname, nil)
	if err != nil {
		return "", err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		result, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(result), nil
	}
}
