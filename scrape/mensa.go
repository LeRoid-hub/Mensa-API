package scrape

import (
	"io"
	"strings"

	"github.com/LeRoid-hub/Mensa-API/models"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeMensa(h io.ReadCloser) models.Mensa {
	var mensa models.Mensa
	var mensaName = ""
	var mensaLocation = ""

	doc, err := goquery.NewDocumentFromReader(h)
	if err != nil {
		return models.Mensa{}
	}

	doc.Find("h1.aw-title-header-title").First().Each(func(i int, s *goquery.Selection) {
		mensaName = s.Text()
	})

	doc.Find("a.panel-body").Each(func(i int, s *goquery.Selection) {
		l, err := s.Html()
		if err != nil {
			return
		}
		l = strings.Replace(l, "<br/>", " ", -1)
		l = strings.Replace(l, "<br>", " ", -1)
		l = strings.Replace(l, "</br>", " ", -1)

		mensaLocation = l
	})

	mensa.SetMensa(mensaName, mensaLocation)

	//Day
	var day models.Day

	doc.Find("h2.aw-menu-title").Each(func(i int, s *goquery.Selection) {
		day.SetDay(s.Text())
	})

	//Menu
	doc.Find("div.aw-meal-category").Each(func(i int, s *goquery.Selection) {

		var menu models.Menu

		s.Find("h3.aw-meal-category-name").Each(func(i int, t *goquery.Selection) {
			menu.SetMenu(t.Text())
		})

		//Meal
		var meal models.Meal

		s.Find("div.aw-meal").Each(func(i int, t *goquery.Selection) {
			mealName := ""
			mealPrice := ""
			mealAttributes := ""
			t.Find("p.aw-meal-description").First().Each(func(i int, u *goquery.Selection) {
				mealName = u.Text()
			})
			t.Find("div.aw-meal-price").First().Each(func(i int, u *goquery.Selection) {
				mealPrice = u.Text()
			})
			t.Find("p.aw-meal-attributes").First().Each(func(i int, u *goquery.Selection) {
				mealAttributes = u.Text()
			})
			meal.SetMeal(mealName, mealPrice, mealAttributes)
			menu.AddMeal(meal)
		})
		day.AddMenu(menu)
	})
	mensa.AddDay(day)

	return mensa

}
