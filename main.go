// School Management System API
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: Oluwapelumi Oyenuga<oluwapelumioyenuga@gmail.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
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
