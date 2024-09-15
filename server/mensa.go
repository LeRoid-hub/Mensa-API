package server

import (
	"github.com/LeRoid-hub/Mensa-API/cache"
	"github.com/LeRoid-hub/Mensa-API/fetch"
	"github.com/LeRoid-hub/Mensa-API/scrape"
	"github.com/gin-gonic/gin"
)

func mensa(c *gin.Context) {
	mensa := c.Param("mensa")
	if mensa == "" {
		c.JSON(400, gin.H{
			"error": "mensa is required",
		})
		return
	}
	city := c.Param("city")
	if city == "" {
		c.JSON(400, gin.H{
			"error": "city is required",
		})
		return
	}

	if cache.HasCacheData(city + "/" + mensa) {
		c.JSON(200, cache.GetCacheData(mensa))
		return
	}

	resp, err := fetch.Fetch(mensa)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"error": "status code is not 200",
		})
		return
	}

	scraped := scrape.ScrapeState(resp.Body)
	c.JSON(200, scraped)
}
