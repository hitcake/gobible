package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		return s
	}
	start := 0
	if s[0] == '+' || s[0] == '-' {
		start = 1
		buf.WriteByte(s[0])
	}
	end := strings.IndexByte(s, '.')
	if end == -1 {
		end = len(s)
	}
	dotBefore := s[start:end]
	dotLen := len(dotBefore)
	if dotLen > 3 {
		for i := dotLen; i > 0; i-- {
			if i%3 == 0 && i != dotLen {
				buf.WriteByte(',')
			}
			buf.WriteByte(dotBefore[dotLen-i])
		}
	} else {
		buf.WriteString(dotBefore)
	}
	buf.WriteString(s[end:])
	return buf.String()
}

func main() {
	fmt.Println(comma("-23456789.12345"))
}
