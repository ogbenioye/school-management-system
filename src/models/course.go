package models

import (
	"school-management-system/src/setup"
	"time"

	"github.com/jinzhu/gorm"
)

type Course struct {
	CourseCode  string    `gorm:"primary_key" json:"code"`
	CourseTitle string    `json:"courseTitle"`
	Students    []Student `gorm:"many2many:student_courses;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
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
	db.Where("Course_Code=?", Id).Find(&course)
	return &course, db
}

func DeleteCourse(Id string) Course {
	var course Course
	db.Where("Course_Code=?", Id).Delete(course)
	return course
}

func ListAllCourses() []Course {
	var Courses []Course
	db.Find(&Courses)
	return Courses
}
