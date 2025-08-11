package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println(WaitForServer("www.jetbrains.com"))
}

func WaitForServer(url string) error {
	const timeout = 30 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.Printf("server not ready yet (%s)", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("timed out waiting for %s", url)
}
