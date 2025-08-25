package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	ch   chan<- string
	name string
}

var (
	entering    = make(chan client)
	leaving     = make(chan client)
	messages    = make(chan string)
	readtimeout = 5 * time.Minute
)

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// 启动一个goroutine去发送消息
			go func() {
				for cli := range clients {
					cli.ch <- msg
				}
			}()
		case cli := <-entering:
			//第一个进来的不提示
			if len(clients) > 0 {
				var buf bytes.Buffer
				buf.WriteString("Current clients :")
				for cli1 := range clients {
					buf.WriteString(cli1.name)
				}
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
		fmt.Fprintf(conn, msg+"\n")
	}
}
func handleConn(conn net.Conn) {
	// 建立缓冲区
	ch := make(chan string, 10)

	go clientWriter(conn, ch)
	ch <- "Please enter your name."
	input := bufio.NewScanner(conn)
	if input.Scan() && input.Text() != "" {
		who := input.Text()
		ch <- "You are " + who
		messages <- who + " has arrived"
		c := client{ch, who}
		entering <- c

		timer := time.NewTimer(readtimeout)
		go func() {
			select {
			case <-timer.C:
				conn.Close()
			}
		}()
		for input.Scan() {
			timer.Reset(readtimeout)
			messages <- who + ": " + input.Text()
		}
		leaving <- c
		messages <- who + " has left"
	} else {
		ch <- "No name provided. You cann't get into it!"
		time.Sleep(1 * time.Second)
	}
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
