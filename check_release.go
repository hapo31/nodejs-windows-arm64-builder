package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
	"time"
)

type Release struct {
	Name				string	`json:"tag_name"`
	Published		string	`json:"published_at"`
}

func main() {
	srcRepoPath := os.Args[1]
	targetRepoPath := os.Args[2]

	if srcRepoPath != "" && targetRepoPath != "" {
		srcURL := fmt.Sprintf("https://api.github.com/repos/%s/releases", srcRepoPath)
		targetURL := fmt.Sprintf("https://api.github.com/repos/%s/releases", targetRepoPath)
		resSrc, _ := http.Get(srcURL)
		resTarget, _ := http.Get(targetURL)

		defer resSrc.Body.Close()
		defer resTarget.Body.Close()

		srcByte, _ := ioutil.ReadAll(resSrc.Body)
		targetByte, _ := ioutil.ReadAll(resTarget.Body)

		var srcRelease []Release
		var targetRelease []Release

		srcErr := json.Unmarshal(srcByte, &srcRelease)
		targetErr := json.Unmarshal(targetByte, &targetRelease)

		if len(srcRelease) == 0 {
			fmt.Print(targetRelease[0].Name)
			os.Exit(0)
		}

		if srcErr != nil || targetErr != nil {
			fmt.Print(srcErr)
			fmt.Print(targetErr)
			os.Exit(1)
		}

		srcRelaseDate, _ := time.Parse("2006-01-02T15:04:05Z", srcRelease[0].Published)
		targetRelaseDate, _ := time.Parse("2006-01-02T15:04:05Z", targetRelease[0].Published)
		
		if targetRelaseDate.Unix() > srcRelaseDate.Unix() {
			fmt.Print(targetRelease[0].Name)
			os.Exit(0)
		} else {
			os.Exit(0)
		}
	}
}