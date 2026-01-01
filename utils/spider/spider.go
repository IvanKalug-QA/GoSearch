package spider

import (
	c "GoSearch/pkg/crawler"
	"GoSearch/utils/file"
	"fmt"
)

func RunSpiner(searchResult string) {
	var crawler c.Crawler = c.CreateCrawler("https://go.dev", "go.dev")
	crawler.Parse()
	output := crawler.GetResult(searchResult)
	if len(output) == 0 {
		fmt.Println("По вашему запросу ничего не было найдено(")
		return
	}
	file.Write(output)
}
