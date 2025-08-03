package main

import (
	"fmt"
	"os"
)

// main method
func main() {
	var str string
	// 遍历os.Args
	for i := 1; i < len(os.Args); i++ {
		str = str + os.Args[i] + " "
	}
	fmt.Println(str)
}
