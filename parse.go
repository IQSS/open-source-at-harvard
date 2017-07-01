package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	tsvFile, _ := os.Create("data.tsv")
	defer tsvFile.Close()
	writer := csv.NewWriter(tsvFile)
	writer.Comma = '\t'
	defer writer.Flush()
	files, _ := filepath.Glob("*.json")
	header := []string{"stars", "language", "updated", "issues", "size", "forks", "watchers", "repo", "created", "description"}
	writer.Write(header)
	for _, element := range files {
		file, e := ioutil.ReadFile(element)
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}
		var repo Repo
		json.Unmarshal(file, &repo)
		fmt.Println(repo.Url, repo.Stars)
		url := repo.Url
		stars := strconv.Itoa(repo.Stars)
		lang := repo.Lang
		openIssues := strconv.Itoa(repo.OpenIssues)
		size := strconv.Itoa(repo.Size)
		forks := strconv.Itoa(repo.Forks)
		watchers := strconv.Itoa(repo.Watchers)
		pushedAt := repo.PushedAt.Format("2006-01-02")
		createdAt := repo.CreatedAt.Format("2006-01-02")
		desc := repo.Desc
		values := []string{stars, lang, pushedAt, openIssues, size, forks, watchers, url, createdAt, desc}
		writer.Write(values)
	}
}

type Repo struct {
	Url        string    `json:"html_url"`
	Stars      int       `json:"stargazers_count"`
	Lang       string    `json:"language"`
	OpenIssues int       `json:"open_issues_count"`
	Size       int       `json:"size"`
	Forks      int       `json:"forks_count"`
	Watchers   int       `json:"watchers_count"`
	PushedAt   time.Time `json:"pushed_at"`
	CreatedAt  time.Time `json:"created_at"`
	Desc       string    `json:"description"`
}
