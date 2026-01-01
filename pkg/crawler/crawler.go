package crawler

import (
	"GoSearch/pkg/index"
	"GoSearch/pkg/search"
	"GoSearch/utils/text"
	"fmt"

	"github.com/gocolly/colly"
)

type Crawler interface {
	Parse()
	CountLinkParsed()
	GetResult(s string) []string
}

type MyCrawler struct {
	url         string
	parserLinks int
	Results     []index.Document
}

func (m *MyCrawler) Parse() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		href := h.Attr("href")

		textWithUrl := text.SimpleNormalize(h.Text)

		fullUrl := h.Request.AbsoluteURL(href)

		if len(textWithUrl) > 1 && len(fullUrl) > 1 {
			m.parserLinks++
			m.Results = append(
				m.Results,
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

func (m *MyCrawler) CountLinkParsed() {
	fmt.Printf("Total link will be parsing: %v", m.parserLinks)
}

func (m *MyCrawler) GetResult(s string) []string {
	indexDoc := index.CreateIndex(m.Results)
	output := search.Search(s, indexDoc, m.Results)
	return output
}

func CreateCrawler(urlParse string) *MyCrawler {
	return &MyCrawler{
		url:     urlParse,
		Results: make([]index.Document, 0),
	}
}
