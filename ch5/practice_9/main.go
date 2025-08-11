package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}

func main() {
	fmt.Println(expand("hello world,foo", strings.ToUpper))
}
