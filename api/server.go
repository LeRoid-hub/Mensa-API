package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/bl/:bundesland", bundesland)
	r.Run("localhost:8080")
}

func bundesland(c *gin.Context) {
	bundesland := c.Param("bundesland")
	if bundesland == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bundesland is required",
		})
		return
	}
	resp, err := Fetch("bl/" + bundesland)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	html := string(d)
	scraped := ScrapeBundesland(html)
	c.JSON(http.StatusOK, scraped)
}

func Fetch(path string) (*http.Response, error) {
	baseurl := "https://www.imensa.de/"
	queryurl := baseurl + "/" + path
	u, err := url.ParseRequestURI(queryurl)
	if err != nil {
		return nil, err
	}
	return http.Get(u.String())
}

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
