package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
bye:
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "cd ") {
			params := strings.Split(scanner.Text(), " ")
			if len(params) != 2 {
				io.WriteString(conn, "command error\n")
				continue
			}
			newdir := dir + string(os.PathSeparator) + params[1]
			_, err := filepath.Abs(newdir)
			if err != nil {
				io.WriteString(conn, params[1]+" dir not found\n")
				continue
			}
			dir = newdir
			io.WriteString(conn, "ok\n")
			continue
		}
		switch scanner.Text() {
		case "pwd":
			_, err = io.WriteString(conn, dir+"\n")
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
		case "list":
			entries, err := os.ReadDir(dir)
			if err != nil {
				fmt.Println("error:", err)
				io.WriteString(conn, "server error:"+err.Error())
			}
			var list []string
			for _, entry := range entries {
				list = append(list, entry.Name())
			}
			_, err = io.WriteString(conn, strings.Join(list, " ")+"\n")
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
		case "quit", "exit", "close":
			_, err := io.WriteString(conn, "bye\n")
			if err != nil {
				fmt.Println("error:", err)
			}
			break bye
		default:
			_, err := io.WriteString(conn, "command not recognized\n")
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
		}
	}
}
