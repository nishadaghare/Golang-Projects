package main

import (
	"MovieDistribution/handlers"
	"MovieDistribution/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	handlers.Regions = services.LoadRegions("cities.csv")

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/distributor", handlers.CreateDistributor)
	router.POST("/permission", handlers.AddPermission)
	router.POST("/check", handlers.CheckPermission)
	router.GET("/regions", func(c *gin.Context) {
		var list []string
		for k := range handlers.Regions {
			list = append(list, k)
		}
		c.JSON(200, list)
	})
	router.Run(":8080")
}
