package router

import (
	"school-management-system/src/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCourseRoutes(router *gin.Engine) {
	router.POST("/sms/course", handler.AddNewCourse)
}
