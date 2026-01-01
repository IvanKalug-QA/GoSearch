package spider

import (
	c "GoSearch/pkg/crawler"
	"GoSearch/utils/file"
	"fmt"
)

const (
	DEFAULT_URL string = "https://go.dev"
	FILE_URLS   string = "urls.txt"
)

func RunSpiner(searchResult string) {
	f := file.ReadUrl(FILE_URLS)

	var crawler c.Crawler

	if f != nil {
		crawler = c.New(f)
		crawler.Parse()
	} else {
		crawler = c.New([]string{DEFAULT_URL})
		crawler.Parse()
	}
	output := crawler.GetResult(searchResult)
	if len(output) == 0 {
		fmt.Println("По вашему запросу ничего не было найдено(")
		return
	}
	file.Write(output)
}
