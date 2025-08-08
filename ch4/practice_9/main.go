package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	worddmap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		text := input.Text()
		worddmap[text]++
	}
	for c, n := range worddmap {
		fmt.Printf("%s\t%d\n", c, n)
	}
}
