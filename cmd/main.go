package main

import (
	c "GoSearch/pkg/crawler"
	"os"
)

func main() {
	ParseArgs()
	if searchParam == "" {
		os.Exit(1)
	}
	var crawler c.Crawler = c.CreateCrawler("https://go.dev", "go.dev")
	crawler.Parse()
	crawler.ShowResult(searchParam)
}
