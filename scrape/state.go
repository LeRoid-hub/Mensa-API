package scrape

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeState(h io.ReadCloser) []string {
	var cities []string

	doc, err := goquery.NewDocumentFromReader(h)
	if err != nil {
		return []string{}
	}

	doc.Find("a.primary").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")

		city := strings.Split(href, "/")[0]

		cities = append(cities, string(city))
	})
	return cities

}
