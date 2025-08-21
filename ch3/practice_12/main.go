package main

import "fmt"

func equal(a, b string) bool {
	aMap := make(map[rune]int)
	bMap := make(map[rune]int)
	for _, v := range a {
		aMap[v]++
	}
	for _, v := range b {
		bMap[v]++
	}
	for k, v := range aMap {
		if bMap[k] != v {
			return false
		}
	}
	for k, v := range bMap {
		if aMap[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(equal("人有所操", "操所有人"))
}
