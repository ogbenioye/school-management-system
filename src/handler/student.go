package handler

import (
	"net/http"
	"school-management-system/src/models"
	"school-management-system/src/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func EnrolStudent(c *gin.Context) {
	student := &models.Student{}

	utils.BindJSON(student, c)
	id := xid.New()
	student.StudentID = "STU/" + strconv.Itoa(id.Time().Year()) + "/" + strconv.Itoa(int(id.Counter()))
	student.EnrolledAt = time.Now()

	s := student.EnrolStudent()

	c.JSON(http.StatusOK, s)
}
