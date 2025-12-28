package crawler

import (
	"GoSearch/pkg/index"
	"GoSearch/utils/text"
	"fmt"

	"github.com/gocolly/colly"
)

type Crawler interface {
	Parse()
	PrintResult()
}

type MyCrawler struct {
	url         string
	urlName     string
	parserLinks int
	results     []index.Document
}

func (m *MyCrawler) Parse() {
	c := colly.NewCollector(colly.AllowedDomains(m.urlName))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		href := h.Attr("href")

		textWithUrl := text.SimpleNormalize(h.Text)

		fullUrl := h.Request.AbsoluteURL(href)

		if len(textWithUrl) > 1 && len(fullUrl) > 1 {
			m.parserLinks++
			m.results = append(
				m.results,
				index.Document{
					ID:    m.parserLinks,
					URL:   fullUrl,
					Title: textWithUrl,
				})
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		// Обработка ошибок
	})
	err := c.Visit(m.url)
	if err != nil {
		panic(err)
	}
}

func (m *MyCrawler) PrintResult() {
	for _, e := range m.results {
		fmt.Println(e)
	}
	fmt.Printf("Total link will be find: %v", m.parserLinks)
}

func CreateCrawler(urlParse, nameUrlParse string) *MyCrawler {
	return &MyCrawler{
		url:     urlParse,
		urlName: nameUrlParse,
		results: make([]index.Document, 0),
	}
}
