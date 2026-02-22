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

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	d, dOk := Distributors[req.Name]
	r, rOk := Regions[req.Region]

	if !dOk {
		c.JSON(400, gin.H{"error": "Distributor not found"})
		return
	}

	if !rOk {
		c.JSON(400, gin.H{"error": "Region not found"})
		return
	}

	allowed := services.HasAccess(d, r)

	c.JSON(200, gin.H{"allowed": allowed})
}
