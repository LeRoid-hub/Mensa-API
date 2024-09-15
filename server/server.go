package server

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/state/:state", state)
	r.GET("/city/:city", city)
	r.GET("/mensa/:city/:mensa", mensa)
	r.Run() // listen and serve on
}