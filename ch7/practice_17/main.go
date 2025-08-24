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

type Selector struct {
	Node string
	Attr map[string]string
}

func (s *Selector) String() string {
	return s.Node
}
func main() {
	data, err := fetch.Get("http://www.w3.org/TR/2006/REC-xml11-20060816")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch err:%s", err)
		return
	}
	selector := make([]*Selector, 0, len(os.Args[1:]))
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "[") && strings.Contains(arg, "]") {
			leftIndex := strings.Index(arg, "[")
			rightIndex := strings.Index(arg, "]")
			if rightIndex < leftIndex || rightIndex != len(arg)-1 {
				fmt.Fprintf(os.Stderr, "selector err:%s arg:%s", arg, arg)
				return
			}
			node := arg[0:leftIndex]
			attrdesc := arg[leftIndex+1 : rightIndex]
			attr := make(map[string]string)
			for _, attrfield := range strings.Split(attrdesc, ",") {
				attrArray := strings.Split(attrfield, "=")
				if len(attrArray) != 2 {
					fmt.Fprintf(os.Stderr, "selector err:%s arg:%s", arg, arg)
					return
				}
				attr[attrArray[0]] = attrArray[1]
			}
			selector = append(selector, &Selector{Node: node, Attr: attr})
		} else {
			selector = append(selector, &Selector{Node: arg, Attr: make(map[string]string)})
		}
	}
	doc := xml.NewDecoder(bytes.NewReader(data))
	var stack []*Selector
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
			attr := make(map[string]string)
			for _, v := range tok.Attr {
				attr[v.Name.Local] = v.Value
			}
			stack = append(stack, &Selector{Node: tok.Name.Local, Attr: attr})
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, selector) {
				for _, sel := range stack {
					fmt.Printf("%s ", sel.Node)
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

func containsAll(x, y []*Selector) bool {
	for len(x) >= len(y) {
		if len(y) == 0 {
			return true
		}
		if x[0].Node == y[0].Node {
			if match(x[0].Attr, y[0].Attr) {
				y = y[1:]
			}

		}
		x = x[1:]
	}
	return false
}

func match(x, y map[string]string) bool {
	for k, v := range y {
		if xv, ok := x[k]; ok && xv == v {
			continue
		}
		return false
	}
	return true
}

/*
./main div div a[title=Markup]
output
html body div div p a : markup
html body div div div p a : Definition
html body div div div p a : markup
html body div div div p a : markup
html body div div div div p a : markup

*/
