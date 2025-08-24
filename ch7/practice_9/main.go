package main

import (
	"fmt"
	"gobible/ch7/practice_8/tracks"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
)

const html = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
</head>
<body>
<table>
<thead>
<th>No</th>
<th><a href="/index?column=0">Title</a></th>
<th><a href="/index?column=1">Artist</a></th>
<th><a href="/index?column=2">Album</a></th>
<th><a href="/index?column=3">Year</a></th>
<th><a href="/index?column=4">Length</a></th>
</thead>
<tbody>
{{range $index, $element := .tracks}}
<tr>
   <td>{{$index}}</td>
	<td>{{$element.Title}}</td>
	<td>{{$element.Artist}}</td>
	<td>{{$element.Album}}</td>
	<td>{{$element.Year}}</td>
    <td>{{$element.Length}}</td>
</tr>
{{end}}
</tbody>
</table>
</body>`

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		column := r.FormValue("column")
		columnIndex := 0
		if column != "" {
			columnIndex, err = strconv.Atoi(column)
			if err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
		}
		sort.Sort(tracks.ByColumn{Tracks: tracks.Tracks, ColumnIndex: columnIndex})
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		template.Must(template.New("escape").Parse(html)).Execute(w, map[string]any{
			"tracks": tracks.Tracks,
		})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
