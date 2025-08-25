package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
)

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			//第一个进来的不提示
			if len(clients) > 0 {
				var buf bytes.Buffer
				buf.WriteString("Current clients :")
				for cli1 := range clients {
					buf.WriteString(cli1.name)
				}
				buf.WriteString("\n")
				cli.ch <- buf.String()
			}

			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)

		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, msg)
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	c := client{ch, who}
	entering <- c

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	leaving <- c
	messages <- who + " has left"
	conn.Close()

}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	go broadcast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
