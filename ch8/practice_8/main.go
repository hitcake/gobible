package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {

	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		c.Close()
	}()
	lines := make(chan string)
	go scan(c, lines)
	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)
	for {
		select {
		case line := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, line, 1*time.Second, &wg)
		case <-timer.C:
			return
		}

	}
}

func scan(r io.Reader, lines chan<- string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Print(err)
	}
}
