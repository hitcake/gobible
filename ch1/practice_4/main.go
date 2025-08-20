package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineFiles := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		//注意 countLines 函数在其声明前被调用。函数和包级别的变量（package-level entities）可以任意顺序声明，
		//并不影响其被调用。（译注：最好还是遵循一定的规范）
		countLines(os.Stdin, counts, lineFiles)
	} else {
		for _, arg := range files {
			// os.Open 函数返回两个值。第一个值是被打开的文件（*os.File），其后被 Scanner 读取。
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, lineFiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, line_files map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			line_files[input.Text()] = line_files[input.Text()] + " " + f.Name()
		}
	}
}

/*
output
practice_4 % go build main.go
practice_4 % ./main ../dup2/dup.txt ../dup2/dup2.txt
2       170      ../dup2/dup.txt
2       181      ../dup2/dup.txt
3       135      ../dup2/dup.txt ../dup2/dup2.txt
3       131      ../dup2/dup.txt ../dup2/dup2.txt
*/
