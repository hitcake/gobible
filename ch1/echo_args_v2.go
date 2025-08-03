package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	for _, arg := range os.Args[1:] {
		str = str + arg + " "
	}
	fmt.Println(str)
}
