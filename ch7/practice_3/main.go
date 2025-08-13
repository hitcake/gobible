package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) add(value int) {
	if t == nil {
		return
	}
	if value < t.value {
		if t.left == nil {
			t.left = &tree{value: value}
		} else {
			t.left.add(value)
		}
	} else {
		if t.right == nil {
			t.right = &tree{value: value}
		} else {
			t.right.add(value)
		}
	}
}

func (t *tree) String() string {
	var buf bytes.Buffer
	if t.left != nil {
		buf.WriteString(t.left.String())
	}
	buf.WriteString(fmt.Sprintf("%d ", t.value))
	if t.right != nil {
		buf.WriteString(t.right.String())
	}
	return buf.String()
}

func main() {
	var root *tree = &tree{value: 1}
	root.add(4)
	root.add(2)
	root.add(3)
	fmt.Println(root.String())
}
