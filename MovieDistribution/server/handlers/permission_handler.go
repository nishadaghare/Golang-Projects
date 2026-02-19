package handlers

import (
	"MovieDistribution/models"
	"MovieDistribution/services"

	"github.com/gin-gonic/gin"
)

var Regions map[string]*models.Region

func CheckPermission(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Region string `json:"region"`
	}

	c.BindJSON(&req)

	d := Distributors[req.Name]
	r := Regions[req.Region]

	allowed := services.HasAccess(d, r)

	c.JSON(200, gin.H{"allowed": allowed})
}
