package main

import (
	"studentAPI/middleware"
	"studentAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(router)
	router.Run("localhost:8080") // listen and serve  "localhost:8080"
}
