package main

import (
	"fmt"
	"log"
)

func add(a, b string) (c string, err error) {
	type d struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case d{}:
			c = "todo"
		default:
			panic(p)
		}
	}()
	panic(d{})
}

func main() {
	result, err := add("hello", "world")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
