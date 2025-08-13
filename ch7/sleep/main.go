package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 5*time.Second, "how long to sleep")

func main() {
	flag.Parse()
	fmt.Printf("sleep for %v\n", *period)
	time.Sleep(*period)
	fmt.Println()
}
