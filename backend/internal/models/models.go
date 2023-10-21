package models

import (
	"time"

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

type Hidden struct{
	User User `gorm:"embedded;"`
	Role string `gorm:"embedded;"`
}

type Teacher struct{
	gorm.Model
	User User `gorm:"embedded;"`
	Pref_auds `gorm:"column:pref_auds" json:"pref_auds"`
	Pref_slots `gorm:"column:pref_slots" json:"pref_slots"`
	WeekendsDay int `gorm:"column:weekends_day" json:"weekends_day"`

}

type Admin struct {
	gorm.Model
	User User `gorm:"embedded;"`
}

type Group struct {
	gorm.Model
	FacultyCode   string `gorm:"column:faculty_code" json:"faculty_code"`
	SchoolClasses int    `gorm:"column:school_classes" json:"school_classes"`
	YearComing    int    `gorm:"column:year_coming" json:"year_coming"`
	Subgroup      int    `gorm:"column:subgroup" json:"subgroup"`
}

type Student struct {
	gorm.Model
	User    User  `gorm:"embedded;"`
	GroupID Group `gorm:"foreignkey:group_id;association_foreignkey:id"`
}

type Subject struct {
	gorm.Model
	Discipline  string  `gorm:"column:discipline" json:"discipline"`
	TeacherID   Teacher `gorm:"foreignkey:teacher_id;association_foreignkey:id"`
	Type        string  `gorm:"column:type" json:"type"`
}

type Audience struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Type string `gorm:"column:type" json:"type"`
}

type Lesson struct{
	gorm.Model
	SubjectID Subject `gorm:"foreignkey:subject_id;association_foreignkey:id"`
	GroupID Group `gorm:"foreignkey:group_id;association_foreignkey:id"`
	LessonHours int     `gorm:"column:lesson_hours" json:"lesson_hours"`
}

type Schedule struct{
	gorm.Model
	LessonID Lesson `gorm:"foreignkey:lesson_id;association_foreignkey:id"`
	AudienceID Audience `gorm:"foreignkey:audience_id;association_foreignkey:id"`
	Day int `gorm:"column:day" json:"day"`
	Time int `gorm:"column:time" json:"time"`
}
