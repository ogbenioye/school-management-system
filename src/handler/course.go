package handler

import (
	"net/http"
	"school-management-system/src/models"
	"school-management-system/src/utils"

	"github.com/gin-gonic/gin"
)

func AddNewCourse(c *gin.Context) {
	course := &models.Course{}

	utils.BindCourse(course, c)

	validateCourse, _ := models.GetCourseByID(course.CourseCode)

	if validateCourse.CourseCode == course.CourseCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": course.CourseCode + " exists already",
		})
		return
	}

	co := course.AddCourse()

	c.JSON(http.StatusOK, co)
}
