package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const BASE_URL = "http://www.omdbapi.com/?t=%s&apikey=%s"

var API_KEY = flag.String("apikey", "", "api key")

type Result struct {
	Title  string
	Poster string
	Year   string
}

func newclient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
	}
}
func saveImg(poster, title string) error {
	resp, err := newclient().Get(poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(title + ".jpeg")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	flag.Parse()
	if len(*API_KEY) == 0 {
		fmt.Fprintln(os.Stderr, "apikey is required")
		os.Exit(1)
	}
	//os.Args[1:] 要改为flag.Args 否则会把-apikey遍历出来
	for _, arg := range flag.Args() {
		url := fmt.Sprintf(BASE_URL, arg, *API_KEY)
		resp, err := newclient().Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Fprintln(os.Stderr, "search query failed:", resp.Status)
			continue
		}
		var result Result
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if saveImg(result.Poster, result.Title); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

	}
}
