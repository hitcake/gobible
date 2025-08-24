package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"gobible/ch5/fetch"
	"io"
	"os"
	"strings"
)

func main() {
	data, err := fetch.Get("http://www.w3.org/TR/2006/REC-xml11-20060816")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch err:%s", err)
		return
	}

	doc := xml.NewDecoder(bytes.NewReader(data))
	var stack []string
	for {
		tok, err := doc.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(x) >= len(y) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
