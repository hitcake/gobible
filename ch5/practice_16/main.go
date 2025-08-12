package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(strings.Join([]string{"hello", "world"}, ","))
	fmt.Println(join(",", "Hello", "world"))
}

func join(sep string, s ...string) string {
	var sb strings.Builder
	l := len(s)
	for i, s := range s {
		sb.WriteString(s)
		if i != l-1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}
