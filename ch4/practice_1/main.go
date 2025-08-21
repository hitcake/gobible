package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(compare(&c1, &c2))
}
func compare(a, b *[32]byte) int {
	var count int
	for i, v := range a {
		if v != b[i] {
			count++
		}
	}
	return count
}
