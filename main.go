package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wuiscmc/ads/controllers"
	"github.com/wuiscmc/ads/repositories"
	"net/http"
)

func main() {

	routes := gin.Default()
	adController := controllers.NewAdController(repositories.NewAdRepository())

	v1 := routes.Group("/v1")
	{
		v1.GET("/ads/:zoneId", func(c *gin.Context) {
			response := adController.FindAd(c.Params.ByName("zoneId"))
			c.JSON(http.StatusOK, response)
		})
	}
	routes.Run(":3000")
}
