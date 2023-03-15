package handler

import (
	"net/http"
	"school-management-system/src/models"
	"school-management-system/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func EnrolStudent(c *gin.Context) {
	var courses []models.Course
	var courseCodes []string
	student := &models.Student{}

	utils.BindStudent(student, c)

	id := xid.New()
	student.StudentID = "STU-" + strconv.Itoa(id.Time().Year()) + "-" + strconv.Itoa(int(id.Counter()))

	//Appending course-codes from the request into a slice
	for i := 0; i < len(student.Courses); i++ {
		courseCodes = append(courseCodes, student.Courses[i].CourseCode)
	}

	//Adding existing course data into the new student being created
	models.DB.Find(&courses, courseCodes)
	student.Courses = courses

	s := student.EnrollStudent()

	c.JSON(http.StatusOK, s)
}

func UpdateStudentInfo(c *gin.Context) {
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

	if studentUpdate.FirstName != "" {
		student.FirstName = studentUpdate.FirstName
	}

	if studentUpdate.LastName != "" {
		student.LastName = studentUpdate.LastName
	}

	if studentUpdate.Gender != "" {
		student.Gender = studentUpdate.Gender
	}

	if studentUpdate.Department != "" {
		student.Department = studentUpdate.Department
	}

	db.Save(student)

	c.JSON(http.StatusOK, student)
}

func ListAllStudents(c *gin.Context) {
	student := models.ListAllStudents()
	c.JSON(http.StatusOK, student)
}

func DisenrollStudent(c *gin.Context) {
	id := c.Param("student-id")

	student, _ := models.GetStudentByID(id)

	models.DisenrollStudent(id)

	if student.StudentID != id {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student disenrolled successfully",
	})
}
