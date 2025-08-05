package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)
	if l <= 3 {
		return s
	}
	for i := l; i > 0; i-- {
		if i%3 == 0 && i != l {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[l-i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("23456789"))
}
