package scrape

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeMensa(h io.ReadCloser) []string {
	var mensas []string

	doc, err := goquery.NewDocumentFromReader(h)
	if err != nil {
		return []string{}
	}

	doc.Find("a.primary").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		print(href)

		mensa := strings.Split(href, "/")[1]

		mensas = append(mensas, string(mensa))
	})
	return mensas

}
