package main

import (
	"school-management-system/src/router"
	"school-management-system/src/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	setup.Connection()

	r := gin.Default()
	router.RegisterStudentRoutes(r)
	router.RegisterCourseRoutes(r)
	router.RegisterSchoolRoutes(r)
	r.Run(":8080")
}
