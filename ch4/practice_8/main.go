package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	digitCounts := make(map[string]int)
	letterCounts := make(map[string]int)
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts[string(r)]++
		} else if unicode.IsDigit(r) {
			digitCounts[string(r)]++
		} else {
			counts[r]++
		}

		utflen[n]++
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range letterCounts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	fmt.Printf("digit\tcount\n")
	for c, n := range digitCounts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("len\tcount\n")
	for c, n := range utflen {
		if n > 0 {
			fmt.Printf("%d\t%d\n", c, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("%d invalid UTF-8 characters\n", invalid)
	}
}
