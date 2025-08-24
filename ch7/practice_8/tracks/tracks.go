package tracks

import (
	"fmt"
	"os"
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

const format = "%v\t%v\t%v\t%v\t%v\t\n"

func PrintTracks(tracks []*Track) {
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

var Tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

type ByColumn struct {
	Tracks      []*Track
	ColumnIndex int
}

func (bc ByColumn) Len() int { return len(bc.Tracks) }
func (bc ByColumn) Less(i, j int) bool {
	switch bc.ColumnIndex {
	case 0:
		return bc.Tracks[i].Title < bc.Tracks[j].Title
	case 1:
		return bc.Tracks[i].Artist < bc.Tracks[j].Artist
	case 2:
		return bc.Tracks[i].Album < bc.Tracks[j].Album
	case 3:
		return bc.Tracks[i].Year < bc.Tracks[j].Year
	case 4:
		return bc.Tracks[i].Length < bc.Tracks[j].Length
	}
	return false
}
func (bc ByColumn) Swap(i, j int) { bc.Tracks[i], bc.Tracks[j] = bc.Tracks[j], bc.Tracks[i] }
