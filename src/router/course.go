package router

import (
	"school-management-system/src/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCourseRoutes(router *gin.Engine) {
	router.POST("/sms/course", handler.AddNewCourse)
	router.PUT("/sms/course/:course-code", handler.UpdateCourseInfo)
	router.GET("/sms/course", handler.ListAllCourses)
	router.DELETE("/sms/course/:course-code", handler.DeleteCourse)
}
