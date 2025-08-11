package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}
func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() {
		log.Printf("exit %s (%s)\n", msg, time.Since(start))
	}
}
