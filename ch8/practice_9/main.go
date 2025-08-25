package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSize)
		} else {
			fileInfo, _ := entry.Info()
			fileSize <- fileInfo.Size()

		}
	}
}

func dirents(dir string) []os.DirEntry {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f MB\n", nfiles, (float64(nbytes))/1024/1024)
}

var verbose = flag.Bool("v", false, "show verbose mode")

var done = make(chan struct{})
var sema = make(chan struct{}, 20)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func caculateSpace(roots []string) {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(time.Second / 2)
	}
	var nfiles, nbytes int64

loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				// do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size

		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func main() {
	tik := time.Tick(3 * time.Second)
	for {
		select {
		case <-tik:
			caculateSpace([]string{"/Users/hit/Downloads"})

		}
	}

}
