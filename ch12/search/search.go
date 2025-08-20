package main

import (
	"fmt"
	"gobible/ch12/params"
	"net/http"
)

type Data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

func search(resp http.ResponseWriter, req *http.Request) {
	var data Data
	data.MaxResults = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resp, "Search: %+v\n", data)
}

func main() {
	http.HandleFunc("/search", search)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
