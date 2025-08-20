package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		defer mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		defer mu.Unlock()
		fmt.Fprintf(w, "Count = %d\n", count)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
