package router

import (
	"school-management-system/src/handler"

	"github.com/gin-gonic/gin"
)

var RegisterStudentRoutes = func(router *gin.Engine) {
	router.POST("/sms/student", handler.EnrolStudent)
	router.PUT("/sms/student/:student-id", handler.UpdateStudentInfo)
	router.GET("/sms/student", handler.ListAllStudents)
	router.DELETE("/sms/student/:student-id", handler.DisenrollStudent)
}
