package main

import "fmt"

type IntList struct {
	value int
	Tail  *IntList
}

func (list *IntList) sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.Tail.sum()
}

func main() {
	var list *IntList
	fmt.Println(list.sum())
}
