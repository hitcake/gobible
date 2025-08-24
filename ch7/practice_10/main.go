package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}
func main() {
	fmt.Println(IsPalindrome(sort.IntSlice{1, 2, 3, 4, 5}))
	fmt.Println(IsPalindrome(sort.IntSlice{1, 2, 3, 2, 1}))
	fmt.Println(IsPalindrome(sort.IntSlice{1, 2, 2, 1}))
}
