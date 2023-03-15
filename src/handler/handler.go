package handler

import (
	"net/http"
	"school-management-system/src/models"

	"github.com/gin-gonic/gin"
)

func ListStudentCourses(c *gin.Context) {
	id := c.Param("student-id")

	student, db := models.GetStudentByID(id)

	db.Preload("Courses").Find(&student)

	if student.StudentID != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found: invalid ID",
		})
		return
	}

	if len(student.Courses) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "no courses registered",
		})
		return
	}
	c.JSON(http.StatusOK, student.Courses)
}

func ListCourseStudents(c *gin.Context) {
	id := c.Param("course-id")

	course, db := models.GetCourseByID(id)

	db.Preload("Students").Find(&course)

	if course.CourseCode != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "course not found: invalid ID",
		})
		return
	}

	if len(course.Students) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "no students registered for this course",
		})
		return
	}

	c.JSON(http.StatusOK, course.Students)
}
