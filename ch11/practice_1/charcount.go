package practice_1

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

type CountResult struct {
	Counts  map[rune]int
	Utflen  []int
	Invalid int
}

func CountCharacters(in *bufio.Reader) (*CountResult, error) {
	results := new(CountResult)
	results.Counts = make(map[rune]int)
	results.Utflen = make([]int, utf8.UTFMax+1)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			results.Invalid++
			continue
		}
		results.Counts[r]++
		results.Utflen[n]++
	}
	return results, nil
}
