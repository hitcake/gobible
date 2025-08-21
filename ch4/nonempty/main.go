package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

/*
*
输入的slice和输出的slice共享一个底层数组。这可以避免分配另一个数组，不过原来的数据将可能会被覆盖
*/
func main() {
	s := []string{"hello", "", "", "world"}
	fmt.Println(nonempty(s)) //[hello world]
	fmt.Println(nonempty(s)) //[hello world world]

}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
