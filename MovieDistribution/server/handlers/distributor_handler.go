package handlers

import (
	"MovieDistribution/models"

	"github.com/gin-gonic/gin"
)

var Distributors = map[string]*models.Distributor{}

func CreateDistributor(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Parent string `json:"parent"`
	}

	c.BindJSON(&req)

	var parent *models.Distributor
	if req.Parent != "" {
		parent = Distributors[req.Parent]
	}

	d := &models.Distributor{
		Name:     req.Name,
		Parent:   parent,
		Includes: map[string]bool{},
		Excludes: map[string]bool{},
	}

	Distributors[req.Name] = d

	c.JSON(200, gin.H{"message": "created"})
}

func AddPermission(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Region string `json:"region"`
	}

	c.BindJSON(&req)

	d := Distributors[req.Name]

	if req.Type == "include" {
		d.Includes[req.Region] = true
	} else {
		d.Excludes[req.Region] = true
	}

	c.JSON(200, gin.H{"message": "updated"})
}
