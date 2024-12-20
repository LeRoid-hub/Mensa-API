package server

import (
	"github.com/LeRoid-hub/Mensa-API/cache"
	"github.com/LeRoid-hub/Mensa-API/fetch"
	"github.com/LeRoid-hub/Mensa-API/scrape"
	"github.com/gin-gonic/gin"
)

func state(c *gin.Context) {
	state := c.Param("state")
	if state == "" {
		c.JSON(400, gin.H{
			"error": "state is required",
		})
		return
	}

	if cache.HasCacheData("state/" + state) {
		cacheData, err := cache.GetCacheData("state/" + state)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, cacheData)
		return
	}

	resp, err := fetch.Fetch(state + ".html")
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
	cache.SetCacheData("state/"+state, scraped)

	c.JSON(200, scraped)
}
