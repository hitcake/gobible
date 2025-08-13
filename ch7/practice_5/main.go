package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type LimitReacher struct {
	reader     io.Reader
	cur, limit int
}

func (lr *LimitReacher) Read(p []byte) (n int, err error) {
	n, err = lr.reader.Read(p[:lr.limit])
	lr.cur += n
	if lr.cur >= lr.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &LimitReacher{reader: r, limit: limit}
}

func main() {
	sr := strings.NewReader("hello world")
	reader := bufio.NewReader(LimitReader(sr, 5))

	for {
		c, err := reader.ReadByte()
		if err == io.EOF {
			break
		} else {
			fmt.Println(string(c))
		}
	}

}
