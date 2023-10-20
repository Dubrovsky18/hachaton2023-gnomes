package services

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
)

type Student interface {
	Get(uuid int) (models.Student, error)
	GetLogin(login string) (models.Student, error)
	Create(student models.Student) error
	Update(student models.Student) error
}

type Teacher interface {
	Create(teacher models.Teacher) error
	Get(uuid int) (models.Teacher, error)
	GetLogin(login string) (models.Teacher, error)
	Update(teacher models.Teacher) error
}

type Admin interface {
	Create(teacher models.Admin) error
	Get(uuid int) (models.Admin, error)
	GetLogin(login string) (models.Admin, error)
	Update(teacher models.Admin) error
}

type Schedule interface {
	Create(schedule []models.Schedule) error
	Get() ([]models.Schedule, error)
	Update(schedule []models.Schedule) error
}

type Service struct {
	Student
	Teacher
	Admin
	Schedule
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Student:  NewStudentService(repos.RepositoryStudent),
		Teacher:  NewTeacherService(repos.RepositoryTeacher),
		Admin:    NewAdminService(repos.RepositoryAdmin),
		Schedule: NewScheduleService(repos.RepositorySchedule),
	}
}
