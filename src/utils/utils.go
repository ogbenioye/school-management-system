package utils

import (
	"net/http"
	"school-management-system/src/models"

	"github.com/gin-gonic/gin"
)

func BindStudent(student *models.Student, c *gin.Context) {
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
}

func BindCourse(course *models.Course, c *gin.Context) {
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
}
