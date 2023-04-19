package models

import (
	"school-management-system/src/setup"
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Student struct {
	StudentID  string   `gorm:"primary_key"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	Gender     string   `json:"gender"`
	Department string   `json:"dept"`
	Courses    []Course `gorm:"many2many:student_courses;" json:"courses"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	setup.Connection()
	DB = setup.GetDB()
	DB.AutoMigrate(&Student{})
}

func (s *Student) EnrollStudent() *Student {
	DB.NewRecord(s)
	DB.Create(&s)
	return s
}

func GetStudentByID(Id string) (*Student, *gorm.DB) {
	var student Student
	DB.Where("Student_ID=?", Id).Find(&student)
	return &student, DB
}

func DisenrollStudent(Id string) Student {
	var student Student
	DB.Where("Student_ID=?", Id).Delete(student)
	return student
}

func ListAllStudents() []Student {
	var Students []Student
	DB.Find(&Students)
	return Students
}
