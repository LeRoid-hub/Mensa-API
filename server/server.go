package server

import (
	"io"
	"mensa/cache"
	"mensa/fetch"
	"mensa/scrape"
	"mensa/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var c = cache.NewCache()

func Run() {
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

	if c.GetCacheData(bundesland) != "" {
		return c.GetCacheData(bundesland)
	}

	resp, err := fetch.Fetch("bl/" + bundesland)
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
	scraped := scrape.ScrapeBundesland(html)
	c.JSON(http.StatusOK, scraped)
}
