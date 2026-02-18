package main

import (
	"./services"

	"./models"

	"github.com/gin-gonic/gin"
)

var regions map[string]*models.Region
var distributors = make(map[string]*models.Distributor)

func main() {
	regions = services.LoadRegions("cities.csv")

	r := gin.Default()

	r.POST("/distributor", createDistributor)
	r.POST("/permission/check", checkPermission)

	r.Run(":8080")
}
