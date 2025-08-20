package main

import (
	"fmt"
	"os"
)

/*
练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
*/
func main() {
	for i, v := range os.Args[1:] {
		fmt.Print(i + 1)
		fmt.Println(" " + v)
	}
}

/*
 practice_2 % go build main.go
 practice_2 % ./main The Go Programming Language
1 The
2 Go
3 Programming
4 Language
 practice_2 %
*/
