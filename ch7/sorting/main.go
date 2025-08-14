package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title, Artist, Album string
	Year                 int
	Length               time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

const format = "%v\t%v\t%v\t%v\t%v\t\n"

func printTracks(tracks []*Track) {
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (a byArtist) Len() int           { return len(a) }
func (a byArtist) Less(i, j int) bool { return a[i].Artist < a[j].Artist }
func (a byArtist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type byYear []*Track

func (a byYear) Len() int           { return len(a) }
func (a byYear) Less(i, j int) bool { return a[i].Year < a[j].Year }
func (a byYear) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type customSort struct {
	t    []*Track
	less func(a, b *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}
func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}
func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func main() {
	// sort.Sort(byArtist(tracks))
	//sort.Sort(byYear(tracks))
	sort.Sort(customSort{tracks, func(a, b *Track) bool {
		if a.Title != b.Title {
			return a.Title < b.Title
		}
		if a.Year != b.Year {
			return a.Year < b.Year
		}
		if a.Length != b.Length {
			return a.Length < b.Length
		}
		return a.Artist < b.Artist
	}})
	printTracks(tracks)

	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // "[4 3 1 1]"
	fmt.Println(sort.IntsAreSorted(values)) // "false
}
