package crawler

import (
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
	results     [][]string
}

func (m *MyCrawler) Parse() {
	c := colly.NewCollector(colly.AllowedDomains(m.urlName))

	c.OnHTML("aside.NavigationDrawer a[href]", func(h *colly.HTMLElement) {
		href := h.Attr("href")

		text := h.Text

		fullUrl := h.Request.AbsoluteURL(href)

		m.parserLinks++

		m.results = append(m.results, []string{fullUrl, text})
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
		fmt.Println("============================")
		fmt.Printf("URL: %v\n", e[0])
		fmt.Printf("TEXT: %v\n", e[1])
		fmt.Println("============================")
	}
}

func CreateCrawler(urlParse, nameUrlParse string) *MyCrawler {
	return &MyCrawler{
		url:     urlParse,
		urlName: nameUrlParse,
		results: make([][]string, 0),
	}
}
