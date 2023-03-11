package models

import (
	"school-management-system/src/setup"
	"time"

	"github.com/jinzhu/gorm"
)

type Student struct {
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	StudentID     string   `gorm:"index:,unique" json:"studentId"`
	Gender        string   `json:"gender"`
	Department    string   `json:"dept"`
	Courses       []Course `gorm:"many2many:student_courses;foreignKey:StudentID" json:"courses"`
	EnrolledAt    time.Time
	LastUpdate    time.Time
	DisenrolledAt *time.Time `sql:"index"`
}

func init() {
	setup.Connection()
	db = setup.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) EnrolStudent() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetStudentByID(Id string) (*Student, *gorm.DB) {
	var student Student
	db.Where("StudentID=?", Id).Find(&student)
	return &student, db
}

func DisenrollStudent(Id string) Student {
	var student Student
	db.Where("StudentID=?", Id).Delete(student)
	return student
}

func ListAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
