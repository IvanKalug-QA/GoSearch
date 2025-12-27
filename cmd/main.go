package main

import (
	c "GoSearch/pkg/crawler"
)

func main() {
	ParseArgs()
	var crawler c.Crawler = c.CreateCrawler("https://go.dev", "go.dev")
	crawler.Parse()
	crawler.PrintResult()
}
