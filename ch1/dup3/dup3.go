package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//ioutil.ReadFile从1.16开始不推荐使用
		//data, err := ioutil.ReadFile(filename)
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/**
output 顺序不确定
dup3 % go build dup3.go
dup3 % ./dup3 ../dup2/dup.txt
2       131
2       135
2       170
2       181
*/
