package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {

	var repoList = "github.tsv"
	tsvFile, tsvOpenError := os.Open(repoList)
	if tsvOpenError != nil {
		os.Exit(1)
	}

	defer tsvFile.Close()

	reader := csv.NewReader(tsvFile)
	reader.Comma = '\t'

	reader.FieldsPerRecord = -1

	_, _ = reader.Read() // delete header

	tsvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each := range tsvData {
		repoUrl := each[0]
		u, err := url.Parse(repoUrl)
		if err != nil {
			panic(err)
		}

		urlPart := strings.Split(u.Path, "/")
		org := urlPart[1]
		repo := urlPart[2]
		orgDashRepo := org + "-" + repo

		gitRepoUrl := "https://api.github.com/repos/" + org + "/" + repo
		fmt.Println("Downloading " + gitRepoUrl)
		response, e := http.Get(gitRepoUrl)
		if e == nil {

			defer response.Body.Close()

			file, err := os.Create(orgDashRepo + ".json")
			if err != nil {
			}
			_, err = io.Copy(file, response.Body)
			if err != nil {
			}
			file.Close()
		} else {
		}
	}

}
