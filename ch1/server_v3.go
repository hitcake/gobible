package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, html.EscapeString(r.URL.Path), r.Proto)
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
		fmt.Fprintf(w, "Host = %q\n", html.EscapeString(r.Host))
		fmt.Fprintf(w, "RemoteAddr = %q\n", html.EscapeString(r.RemoteAddr))
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
