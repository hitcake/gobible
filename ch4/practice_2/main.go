package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"golang.org/x/crypto/sha3"
	"os"
)

var sha = flag.String("sha", "sha256", "SHA")

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		switch *sha {
		case "sha256":
			fmt.Println(sha256.Sum256([]byte(scanner.Text())))
		case "sha512":
			fmt.Println(sha512.Sum512([]byte(scanner.Text())))
		case "sha384":
			fmt.Println(sha3.Sum384([]byte(scanner.Text())))
		default:
			fmt.Println("Unknown SHA")
		}
	}

}
