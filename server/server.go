package server

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Mensen API",
		})
	})
	r.GET("/state/:state", state)
	r.GET("/city/:city", city)
	r.GET("/mensa/:city/:mensa", mensa)
	r.Run(":80")
}
