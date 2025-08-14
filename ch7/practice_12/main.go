package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

const tmpl = `<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8">
	<title>价格表</title>
	</head>
	<body>
	<ul>
        {{range $key, $value := .data}}
            <li>{{$key}}: {{$value}}</li>
        {{end}}
    </ul>
	</body>
	`

func (db database) list(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("list").Parse(tmpl)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	err = t.Execute(w, map[string]interface{}{"data": db})
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	//for item, price := range db {
	//	fmt.Fprintf(w, "%s: %s\n", item, price)
	//}
}
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
func (db database) praseParamater(w http.ResponseWriter, req *http.Request) (string, dollars, error) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameters error\n")
		return "", (dollars(0)), fmt.Errorf("parameter error\n")
	}
	priceF, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter price error\n")
		return "", (dollars(0)), fmt.Errorf("parameter price error\n")
	}
	return item, dollars(priceF), nil
}
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item, price, err := db.praseParamater(w, req)
	if err != nil {
		return
	}
	db[item] = price
	fmt.Fprintf(w, "creaet successfully\n")
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item, price, err := db.praseParamater(w, req)
	if err != nil {
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	db[item] = price
	fmt.Fprintf(w, "updated successfully\n")
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "parameter error\n")
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "deleted successfully\n")
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
