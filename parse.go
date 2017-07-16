package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
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
	var repos []Repo
	for _, element := range files {
		file, e := ioutil.ReadFile(element)
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}
		var repo Repo
		json.Unmarshal(file, &repo)
		fmt.Println(repo.Url, repo.Stars)
		repos = append(repos, repo)
	}
	sort.Sort(ByStars(repos))
	header := []string{"stars", "language", "updated", "issues", "size", "forks", "watchers", "repo", "created", "description"}
	writer.Write(header)
	for _, repo := range repos {
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

type ByStars []Repo

func (a ByStars) Len() int           { return len(a) }
func (a ByStars) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStars) Less(j, i int) bool { return a[i].Stars < a[j].Stars }
