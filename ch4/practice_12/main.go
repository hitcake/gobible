package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Cartoon struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

const baseUrl = "https://xkcd.com/"

func buildIndexFile(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return doBuildIndexFile(filePath)
		}
		return err
	}
	return nil
}

func doBuildIndexFile(filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	write := csv.NewWriter(out)
	write.Comma = '\t'
	write.Write([]string{"title", "img"})
	for i := 1; i < 100; i++ {
		cartoon := &Cartoon{}
		url := fmt.Sprintf("%s%d/info.0.json", baseUrl, i)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			continue
		}
		if err := json.NewDecoder(resp.Body).Decode(cartoon); err != nil {
			resp.Body.Close()
			fmt.Println(err)
			continue
		}
		resp.Body.Close()
		write.Write([]string{cartoon.Title, cartoon.Img})
		write.Flush()
	}
	return nil
}

func search(filePath, name string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	csvr := csv.NewReader(f)
	csvr.Comma = '\t'
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		if record[0] == name {
			return record[1], nil
		}
	}
	return "", fmt.Errorf("%s not found", name)
}

func main() {
	indexFile := "/tmp/index.csv"
	err := buildIndexFile(indexFile)
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range os.Args[1:] {
		img, err := search(indexFile, val)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(img)
		}
	}
}
