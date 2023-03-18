package router

import (
	"school-management-system/src/handler"

	"github.com/gin-gonic/gin"
)

func RegisterSchoolRoutes(router *gin.Engine) {
	router.GET("/sms/:student-id", handler.ListStudentCourses)
	router.GET("/sms/c/:course-id", handler.ListCourseStudents)
	router.PUT("/sms/:student-id", handler.AddToStudentCourses)
	router.PUT("/sms/s/:student-id", handler.RemoveStudentCourses)
}
