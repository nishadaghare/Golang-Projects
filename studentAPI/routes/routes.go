package routes

import (
	"studentAPI/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/students", controller.GetStudents)
	router.GET("/students/:id", controller.GetStudentByID)
	router.POST("/student", controller.CreateStudent)
	router.PUT("/student/:id", controller.UpdateStudent)
	router.DELETE("/student/:id", controller.DeleteStudent)
	router.DELETE("/deletestudents", controller.DeleteAllStudents) // optional

}
