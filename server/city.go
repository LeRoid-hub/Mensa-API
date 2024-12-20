package server

import (
	"github.com/LeRoid-hub/Mensa-API/cache"
	"github.com/LeRoid-hub/Mensa-API/fetch"
	"github.com/LeRoid-hub/Mensa-API/scrape"
	"github.com/gin-gonic/gin"
)

func city(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		c.JSON(400, gin.H{
			"error": "city is required",
		})
		return
	}

	if cache.HasCacheData("city/" + city) {
		cacheData, err := cache.GetCacheData("city/" + city)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, cacheData)
		return
	}

	resp, err := fetch.Fetch(city)
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

	cache.SetCacheData("city/"+city, scraped)
	c.JSON(200, scraped)
}
