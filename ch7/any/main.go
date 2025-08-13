package main

import "fmt"

func main() {
	var any interface{}
	any = 42
	fmt.Printf("%v", any)
	any = 12.34
	fmt.Printf("%v", any)
	any = "hello"
	fmt.Printf("%v", any)
	any = make(chan int)
	fmt.Printf("%v", any)
	any = make(map[int]int)
	fmt.Printf("%v", any)
}
