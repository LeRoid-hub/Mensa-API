package main

import (
	"net/http"
	"net/url"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/fetch/:url", fetchRoute)
	r.Run("localhost:8080")
}

func fetchRoute(c *gin.Context) {
	url := c.Param("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url is required",
		})
		return
	}
	resp, err := Fetch(url)
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
	d = string(d)
	c.JSON(http.StatusOK, d)

}

func Fetch(path string) (*http.Response, error) {
	baseurl := "https://www.imensa.de/"
	queryurl  := baseurl + "/" + path
	u, err := url.ParseRequestURI(queryurl)
	if err != nil {
		return nil, err
	}

	return http.Get(u.String())

}
