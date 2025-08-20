package main

import "fmt"

func main() {
	fmt.Println(PopCount(5))
}

func PopCount(x uint64) int {
	var count int
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
