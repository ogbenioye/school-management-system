package handler

import (
	"net/http"
	"school-management-system/src/models"
	"school-management-system/src/utils"

	"github.com/gin-gonic/gin"
)

type studentCourse struct {
	studentId  string
	courseCode string
}

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

func AddToStudentCourses(c *gin.Context) {
	var courses []models.Course
	courseMap := make(map[string]string)
	var courseCodes []string

	id := c.Param("student-id")

	student, db := models.GetStudentByID(id)

	var studentUpdate models.Student

	utils.BindStudent(&studentUpdate, c)

	if student.StudentID != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	if len(studentUpdate.Courses) == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you have not chosen a course to register",
		})
		return
	}

	if student.Courses != nil {
		for i := 0; i < len(student.Courses); i++ {
			courseMap[student.Courses[i].CourseCode] = student.Courses[i].CourseCode
		}
	}

	for i := 0; i < len(studentUpdate.Courses); i++ {
		courseMap[studentUpdate.Courses[i].CourseCode] = studentUpdate.Courses[i].CourseCode
	}

	for _, v := range courseMap {
		courseCodes = append(courseCodes, v)
	}

	models.DB.Find(&courses, courseCodes)
	student.Courses = courses

	db.Save(student)
	c.JSON(http.StatusOK, student)
}

func RemoveStudentCourses(c *gin.Context) {
	id := c.Param("student-id")

	student, db := models.GetStudentByID(id)

	var studentUpdate models.Student

	utils.BindStudent(&studentUpdate, c)

	if student.StudentID != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	if len(studentUpdate.Courses) == 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you have not chosen a course to unregister",
		})
		return
	}

	var count uint64

	for i := 0; i < len(studentUpdate.Courses); i++ {
		db.Table("student_courses").Where("student_student_id=? AND course_course_code=?", id, studentUpdate.Courses[i].CourseCode).Count(&count)

		if count > 0 {
			db.Table("student_courses").Where("student_student_id=? AND course_course_code=?", id, studentUpdate.Courses[i].CourseCode).Delete(&studentCourse{})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "cannot delete. this student has not registered " + studentUpdate.Courses[i].CourseCode + " prior",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Course removed successfully",
	})
}
