package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Issue struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Number int    `json:"number"`
}

func (s Issue) String() string {
	return fmt.Sprintf("#%d %s #%s", s.Number, s.Title, s.Body)
}

const baseURL = "https://api.github.com/repos"

func newclient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5, // 设置超时时间为10秒
	}
}
func GetIssues(owner, repo string) ([]Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues", baseURL, owner, repo)
	fmt.Printf("GET %s\n", url)
	resp, err := newclient().Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func GetIssue(owner, repo string, number int) (*Issue, error) {
	url := fmt.Sprintf("%s/%s/%s/issues/%d", baseURL, owner, repo, number)
	resp, err := newclient().Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func CreateIssue(owner, repo, token string, issue *Issue) error {
	url := fmt.Sprintf("%s/%s/%s/issues?access_token=%s", baseURL, owner, repo, token)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(*issue); err != nil {
		return err
	}
	resp, err := newclient().Post(url, "application/json", &buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(issue); err != nil {
		return err
	}
	return nil
}

func EditIssue(owner, repo, token string, number int, issue *Issue) error {
	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(*issue); err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s/%s/issues/%d?access_token=%s", baseURL, owner, repo, number, token)
	req, err := http.NewRequest("PATCH", url, &buffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := newclient().Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(issue); err != nil {
		return err
	}
	return nil
}

func main() {
	issues, err := GetIssues("hitcake", "gobible")
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range issues {
		fmt.Println(issue)
	}
	//issue := Issue{Title: "测试",Body: "这是一个API创建的测试"}
	//if CreateIssue("hitcake","gobible","")
}
