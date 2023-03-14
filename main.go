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
	r.Run(":8080")
}
