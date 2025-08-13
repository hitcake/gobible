package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordsAndLinesCounter struct {
	wc int
	lc int
}

func counter(p []byte, split bufio.SplitFunc) (int, error) {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(split)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Scanner.Scan: %w", err)
	}
	return count, nil
}
func (walc *WordsAndLinesCounter) Write(p []byte) (int, error) {
	count, err := counter(p, bufio.ScanWords)
	if err != nil {
		return 0, fmt.Errorf("counter for words: %w", err)
	}
	(*walc).wc += count
	count, err = counter(p, bufio.ScanLines)
	if err != nil {
		return 0, fmt.Errorf("counter for lines: %w", err)
	}
	(*walc).lc += count
	return len(p), nil
}
func (walc WordsAndLinesCounter) String() string {
	return fmt.Sprintf("Words: %d, Lines: %d", walc.wc, walc.lc)
}
func main() {
	var wc WordsAndLinesCounter
	fmt.Fprintf(&wc,
		`There are still some outstanding problems.
				A lot of work is still outstanding.
				At last we could see the dim outline of an island.`)
	fmt.Println(wc)
}
