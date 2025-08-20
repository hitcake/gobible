package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: fetch URL filename")
		return
	}
	url := os.Args[1]
	filename := os.Args[2]
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file failed: %v\n", err)
		return
	}
	start := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(start).Seconds()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}

	b, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get absolute path failed: %v\n", err)
	} else {
		fmt.Printf("%d bytes written to %s\n fetch %s used %.2f", b, absPath, url, secs)
	}

	file.Close()
}
