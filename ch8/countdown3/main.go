package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			abort <- struct{}{}
		}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launching")
}
