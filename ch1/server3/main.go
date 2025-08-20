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

/*
output
GET / HTTP/1.1
Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,*\/*;q=0.8"]
Header["Accept-Language"] = ["zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"]
Header["Upgrade-Insecure-Requests"] = ["1"]
Header["Sec-Fetch-Dest"] = ["document"]
Header["Sec-Fetch-Mode"] = ["navigate"]
Header["Sec-Fetch-User"] = ["?1"]
Header["Priority"] = ["u=0, i"]
Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:141.0) Gecko/20100101 Firefox/141.0"]
Header["Accept-Encoding"] = ["gzip, deflate, br, zstd"]
Header["Connection"] = ["keep-alive"]
Header["Sec-Fetch-Site"] = ["none"]
Host = "localhost:8080"
RemoteAddr = "127.0.0.1:55328"
*/
