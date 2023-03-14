package router

import (
	"school-management-system/src/handler"

	"github.com/gin-gonic/gin"
)

var RegisterStudentRoutes = func(router *gin.Engine) {
	router.POST("/sms", handler.EnrolStudent)
	router.PUT("/sms/:student-id", handler.UpdateStudentInfo)
	router.GET("/sms", handler.ListAllStudents)
	router.DELETE("/sms/:student-id", handler.DisenrollStudent)
}
