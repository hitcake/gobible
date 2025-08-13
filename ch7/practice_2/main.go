package main

import (
	"bytes"
	"io"
	"log"
)

type WriterWrapper struct {
	writer io.Writer
	length int64
}

func (wrapper *WriterWrapper) Write(p []byte) (n int, err error) {
	n, err = wrapper.writer.Write(p)
	wrapper.length += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapper := WriterWrapper{w, 0}
	return &wrapper, &(wrapper.length)
}

func main() {
	b := &bytes.Buffer{}
	c, n := CountingWriter(b)
	data := []byte("hello,world")
	_, err := c.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	if *n != int64(len(data)) {
		log.Fatalf("expected %d, got %d", len(data), *n)
	} else {
		log.Printf("success n=%d", *n)
	}
}
