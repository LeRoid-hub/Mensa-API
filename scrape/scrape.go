package scrape

import (
	"strings"

	"golang.org/x/net/html"
)

func ScrapeBundesland(h string) []string {
	tkn := html.NewTokenizer(strings.NewReader(h))

	var mensen []string

	for {
		if tkn.Next() == html.ErrorToken {
			return mensen
		}

		t := tkn.Token()
		attr := t.Attr

		for _, a := range attr {
			if a.Key == "class" && a.Val == "elements" {
				print(t.Data)
				mensen = append(mensen, t.Data)
			}
		}
	}
}
