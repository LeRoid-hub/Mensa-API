package scrape

import (
	"io"
	"strings"

	"github.com/LeRoid-hub/Mensa-API/models"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeMensa(h io.ReadCloser) models.Mensa {
	var mensa models.Mensa

	doc, err := goquery.NewDocumentFromReader(h)
	if err != nil {
		return models.Mensa{}
	}

	doc.Find("h1.aw-title-header-title").First().Each(func(i int, s *goquery.Selection) {
		mensa.Name = s.Text()
	})

	doc.Find("a.panel-body").Each(func(i int, s *goquery.Selection) {
		l, err := s.Html()
		if err != nil {
			return
		}
		l = strings.Replace(l, "<br/>", " ", -1)
		l = strings.Replace(l, "<br>", " ", -1)
		l = strings.Replace(l, "</br>", " ", -1)

		mensa.Location = l
	})

	//Day
	var day models.Day

	doc.Find("h2.aw-menu-title").Each(func(i int, s *goquery.Selection) {
		day.DayName = s.Text()
	})

	//Menu
	var menu models.Menu

	doc.Find("div.aw-meal-category").Each(func(i int, s *goquery.Selection) {

		s.Find("h3.aw-meal-category-name").Each(func(i int, t *goquery.Selection) {
			menu.Name = t.Text()
		})

		//Meal
		var meal models.Meal

		s.Find("div.aw-meal").Each(func(i int, t *goquery.Selection) {
			t.Find("p.aw-meal-description").First().Each(func(i int, u *goquery.Selection) {
				meal.Name = u.Text()
			})
			t.Find("div.aw-meal-price").First().Each(func(i int, u *goquery.Selection) {
				meal.Price = u.Text()
			})
			t.Find("p.aw-meal-attributes").First().Each(func(i int, u *goquery.Selection) {
				meal.Attributes = u.Text()
			})
			menu.Meal = append(menu.Meal, meal)
		})
		day.Menu = append(day.Menu, menu)
	})
	mensa.AddDay(day)

	return mensa

}
