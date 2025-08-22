package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func removeSpaces(b []byte) []byte {
	for i := 0; i < len(b); {
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(second) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
				continue
			}

		}
		i += size
	}
	return b
}

func main() {
	fmt.Println(string(removeSpaces([]byte("我是   真的    真的  用了心"))))
}
