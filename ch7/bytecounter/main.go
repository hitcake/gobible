package main

import "fmt"

type byteCounter int

func (bc *byteCounter) Write(p []byte) (int, error) {
	*bc += byteCounter(len(p))
	return len(p), nil
}

func main() {
	var bc byteCounter
	bc.Write([]byte("hello"))
	fmt.Println(bc)
	bc = 0
	var name = "Dolly"
	fmt.Fprintf(&bc, "hello, %s", name)
	fmt.Println(bc)
}
