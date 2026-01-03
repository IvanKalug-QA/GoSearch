package spider

import (
	c "GoSearch/pkg/crawler"
	"GoSearch/pkg/netsrv"
	"GoSearch/utils/file"
	"GoSearch/utils/text"
)

const (
	DEFAULT_URL string = "https://go.dev"
	FILE_URLS   string = "urls.txt"
)

func RunSpiner(searchResult string) {
	f := file.ReadUrl(FILE_URLS)
	var searchUrls []string

	if f != nil {
		searchUrls = f
	} else {
		searchUrls = []string{DEFAULT_URL}
	}

	var crawler c.Crawler

	crawler = c.New(searchUrls)
	crawler.Parse()

	if searchResult != "" {
		output := crawler.GetResult(searchResult)
		if text.CheckResult(output) {
			file.Write(output)
		}
	} else {
		netsrv.StartServer(crawler)
	}
}
