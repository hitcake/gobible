package main

import "fmt"

func main() {
	fmt.Println(PopCount2(5))
}

func PopCount2(x uint64) int {
	var count int
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
