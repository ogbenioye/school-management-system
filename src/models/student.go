package models

import (
	"school-management-system/src/setup"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

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
	db = setup.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) EnrollStudent() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetStudentByID(Id string) (*Student, *gorm.DB) {
	var student Student
	db.Where("Student_ID=?", Id).Find(&student)
	return &student, db
}

func DisenrollStudent(Id string) Student {
	var student Student
	db.Where("Student_ID=?", Id).Delete(student)
	return student
}

func ListAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
