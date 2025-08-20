// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

// main method
func main() {
	var s, sep string
	// 遍历os.Args
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
