package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//注意 countLines 函数在其声明前被调用。函数和包级别的变量（package-level entities）可以任意顺序声明，
		//并不影响其被调用。（译注：最好还是遵循一定的规范）
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open 函数返回两个值。第一个值是被打开的文件（*os.File），其后被 Scanner 读取。
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

/*
output
 dup2 % go build dup2.go
 dup2 % ./dup2 dup.txt
2       135
2       131
2       170
2       181
*/
