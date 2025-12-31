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
	PrintParsed()
	ShowResult(s string)
}

type MyCrawler struct {
	url         string
	urlName     string
	parserLinks int
	Results     []index.Document
}

func (m *MyCrawler) Parse() {
	c := colly.NewCollector(colly.AllowedDomains(m.urlName))

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

func (m *MyCrawler) PrintParsed() {
	for _, e := range m.Results {
		fmt.Println(e)
	}
	fmt.Printf("Total link will be find: %v", m.parserLinks)
}

func (m *MyCrawler) ShowResult(s string) {
	indexDoc := index.CreateIndex(m.Results)
	output := search.Search(s, indexDoc, m.Results)
	for _, o := range output {
		fmt.Println(o)
	}
}

func CreateCrawler(urlParse, nameUrlParse string) *MyCrawler {
	return &MyCrawler{
		url:     urlParse,
		urlName: nameUrlParse,
		Results: make([]index.Document, 0),
	}
}
