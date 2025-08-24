package main

import (
	"bufio"
	"fmt"
	"gobible/ch7/practice_8/tracks"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("sort index 0 Title 1 Artist 2 Album 3 Year 4 Lenght\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		columnIndex, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Println("input error, not a number")
		}
		if columnIndex < 0 || columnIndex > 4 {
			fmt.Println("column index out of range, between 0 and 4")
		}
		sort.Sort(tracks.ByColumn{Tracks: tracks.Tracks, ColumnIndex: int(columnIndex)})
		tracks.PrintTracks(tracks.Tracks)
	}
}
