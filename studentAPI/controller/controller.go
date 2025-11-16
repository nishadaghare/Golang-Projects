package controller

import (
	"net/http"
	"strconv"
	"studentAPI/model"

	"github.com/gin-gonic/gin"
)

var Students = []model.Student{}
var AutoIncrementID = 1

// CreateStudent
func CreateStudent(c *gin.Context) {
	var input model.Student

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = AutoIncrementID
	AutoIncrementID++

	Students = append(Students, input)

	c.JSON(http.StatusCreated, gin.H{"data": input})
}

// GetStudents
func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Students})
}

// GetStudentByID
func GetStudentByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, s := range Students {
		if s.ID == id {
			c.JSON(http.StatusOK, gin.H{"data": s})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// UpdateStudent
func UpdateStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input model.Student

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, s := range Students {
		if s.ID == id {
			Students[i].Name = input.Name
			Students[i].Email = input.Email
			Students[i].Age = input.Age

			c.JSON(http.StatusOK, gin.H{"data": Students[i]})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// DeleteStudent
func DeleteStudent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, s := range Students {
		if s.ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

// DeleteAllStudents
func DeleteAllStudents(c *gin.Context) {
	Students = []model.Student{}
	AutoIncrementID = 1

	c.JSON(http.StatusOK, gin.H{"message": "All students deleted"})
}
