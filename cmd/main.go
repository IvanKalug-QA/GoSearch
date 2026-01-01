package main

import (
	c "GoSearch/pkg/crawler"
	"GoSearch/utils/text"
	"fmt"
	"os"
)

func main() {
	ParseArgs()
	if searchParam == "" {
		os.Exit(1)
	}
	var crawler c.Crawler = c.CreateCrawler("https://go.dev", "go.dev")
	crawler.Parse()
	output := crawler.GetResult(searchParam)
	if len(output) == 0 {
		fmt.Println("По вашему запросу ничего не было найдено(")
	}
	text.WriteFile(output)
}
