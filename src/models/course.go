package models

import (
	"school-management-system/src/setup"

	"github.com/jinzhu/gorm"
)

// var db *gorm.DB

type Course struct {
	CourseCode  string    `gorm:"primary_key"`
	CourseTitle string    `json:"courseTitle"`
	Students    []Student `gorm:"many2many:student_courses;"`
}

func init() {
	setup.Connection()
	db = setup.GetDB()
	db.AutoMigrate(&Course{})
}

func (c *Course) AddCourse() *Course {
	db.NewRecord(c)
	db.Create(c)
	return c
}

func GetCourseByID(Id string) (*Course, *gorm.DB) {
	var course Course
	db.Where("CourseCode=?", Id).Find(&course)
	return &course, db
}

func DeleteCourse(Id string) Course {
	var course Course
	db.Where("CourseCode=?", Id).Delete(course)
	return course
}

func ListAllCourses() []Course {
	var Courses []Course
	db.Find(&Courses)
	return Courses
}
