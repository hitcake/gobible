package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var ts = make(map[string]net.Conn)

func showTime() {
	for {
		for k, v := range ts {
			scanner := bufio.NewScanner(v)
			if scanner.Scan() {
				fmt.Println("+-Address-+------time------+")
				fmt.Printf("%s : %s\n", k, scanner.Text())
			}
		}
		fmt.Printf("\n\n\n")
	}
}

func main() {
	//ts := make(map[string][]string)

	for _, arg := range os.Args[1:] {
		split := strings.Split(arg, "=")
		conn, err := net.Dial("tcp", split[1])
		if err != nil {
			log.Print(err)
			continue
		}
		ts[split[0]] = conn
	}
	showTime()

}
