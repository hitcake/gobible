package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"zh-CN"}}
	m.Add("format", "json")
	m.Add("value", "Hello, world!")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("format"))
	fmt.Println(m["value"])

	m = nil
	fmt.Println(m.Get("lang"))
	m.Add("format", "json")
}
