package main

import (
	"GoBible/ch4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	mbefore := time.Now().AddDate(0, -1, 0)
	yBefore := time.Now().AddDate(-1, 0, 0)

	var inAMonth = []*github.Issue{}
	var inAYear = []*github.Issue{}
	var outAYear = []*github.Issue{}
	for _, item := range result.Items {
		if item.CreatedAt.After(mbefore) {
			inAMonth = append(inAMonth, item)
		} else if item.CreatedAt.After(yBefore) {
			inAYear = append(inAYear, item)
		} else {
			outAYear = append(outAYear, item)
		}
	}
	fmt.Printf("%d issues in a month ----------:\n", len(inAMonth))
	for _, item := range inAMonth {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Printf("%d issues in a year ----------:\n", len(inAYear))
	for _, item := range inAYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Printf("%d issues before a year ----------:\n", len(outAYear))
	for _, item := range outAYear {
		fmt.Printf("#%-5d %9.9s %.55s %s\n", item.Number, item.User.Login, item.Title, item.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
