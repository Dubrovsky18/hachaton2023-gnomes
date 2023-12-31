package models

import (
	"gorm.io/gorm"
)

// Column names at database like "column_name"
// SecondName: "second_name"
type User struct {
	Name       string `gorm:"column:name" json:"name"`
	SecondName string `gorm:"column:second_name" json:"second_name"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Email      string `gorm:"column:email" json:"email"`
	Password   string `gorm:"column:password" json:"password"`
}

type Hidden struct {
	User User   `gorm:"embedded;"`
	Role string `gorm:"embedded;"`
}

type Teacher struct {
	gorm.Model
	User        User `gorm:"embedded;"`
	Hours       int  `gorm:"column:hours" json:"hours"`
	WeekendsDay int  `gorm:"column:weekends_day" json:"weekends_day"`
}

type Admin struct {
	gorm.Model
	User User `gorm:"embedded;"`
}

type Group struct {
	gorm.Model
	FacultyCode   string `gorm:"faculty_code" json:"faculty_code"`
	SchoolClasses int    `gorm:"school_classes" json:"school_classes"`
	YearComing    int    `gorm:"year_coming" json:"year_coming"`
	Subgroup      int    `gorm:"subgroup" json:"subgroup"`
}

type Student struct {
	gorm.Model
	User    User  `gorm:"embedded;"`
	GroupID Group `gorm:"group_id;"`
}

type Subject struct {
	gorm.Model
	Discipline string  `gorm:"discipline" json:"discipline"`
	TeacherID  Teacher `gorm:"teacher_id;"`
	Type       string  `gorm:"column:type" json:"type"`
}

type Audience struct {
	gorm.Model
	Name string `gorm:"name" json:"name"`
	Type string `gorm:"type" json:"type"`
}

type Lesson struct {
	gorm.Model
	SubjectID   Subject `gorm:"subject_id"`
	GroupID     Group   `gorm:"group_id"`
	LessonHours int     `gorm:"column:lesson_hours" json:"lesson_hours"`
}

type Schedule struct {
	gorm.Model
	LessonID   Lesson   `gorm:"lesson_id"`
	AudienceID Audience `gorm:"audience_id" json:"audience_id"`
	Day        int      `gorm:"day" json:"day"`
	Time       int      `gorm:"time" json:"time"`
}
