package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Copy() IntSet {
	var cp IntSet
	cp.words = make([]uint64, len(s.words))
	copy(cp.words, s.words)
	return cp
}

func (s *IntSet) remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] = s.words[word] ^ (1 << bit)
	}
}

func (s *IntSet) Clear() {
	for i := 0; i < len(s.words); i++ {
		s.words[i] = 0
	}
}

func (s *IntSet) Len() int {
	var l = 0
	for _, word := range s.words {
		if word != 0 {
			for j := 0; j < 64; j++ {
				if word&(1<<uint(j)) != 0 {
					l++
				}
			}
		}
	}
	return l
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(64)
	x.Add(167)
	fmt.Println(x.Len())
	x.remove(64)
	fmt.Println(x.String())
	y := x.Copy()
	fmt.Println(y.String())
	y.Clear()
	fmt.Println(y.Len())

}
