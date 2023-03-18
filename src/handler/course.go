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

func UpdateCourseInfo(c *gin.Context) {
	//Note: Cannot update course-code because it is a primary key.
	//Any case where it'll be needed to update course-code, you'll need to delete and
	//recreate that course with the accurate/updated coursecode
	id := c.Param("course-code")

	course, db := models.GetCourseByID(id)

	var courseUpdate models.Course

	utils.BindCourse(&courseUpdate, c)

	if course.CourseCode != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "course not found",
		})
		return
	}

	if courseUpdate.CourseCode != "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "cannot update course-code field, delete course and re-register course with correct information",
		})
		return
	}

	if courseUpdate.CourseTitle != "" {
		course.CourseTitle = courseUpdate.CourseTitle
	}

	db.Save(course)

	c.JSON(http.StatusOK, course)
}

func ListAllCourses(c *gin.Context) {
	course := models.ListAllCourses()
	c.JSON(http.StatusOK, course)
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("course-code")

	course, _ := models.GetCourseByID(id)

	if course.CourseCode != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "course not found",
		})
		return
	}

	models.DeleteCourse(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "course deleted successfully",
	})
}
