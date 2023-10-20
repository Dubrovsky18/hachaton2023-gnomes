package repository

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository/postgresql"
	"gorm.io/gorm"
)

type RepositoryStudent interface {
	Create(student models.Student) error
	Get(uuid int) (models.Student, error)
	GetLogin(login string) (models.Student, error)
	Update(student models.Student) error
}

type RepositoryTeacher interface {
	Create(teacher models.Teacher) error
	Get(uuid int) (models.Teacher, error)
	GetLogin(login string) (models.Student, error)
	Update(teacher models.Teacher) error
}

type RepositoryAdmin interface {
	Create(admin models.Admin) error
	Get(uuid int) (models.Admin, error)
	GetLogin(login string) (models.Student, error)
	Update(admin models.Admin) error
}

type RepositorySchedule interface {
	Create(schedule []models.Schedule) error
	Get(uuid int) ([]models.Schedule, error)
	Update(schedule []models.Schedule) error
}

type Repository struct {
	RepositoryStudent
	RepositoryTeacher
	RepositoryAdmin
	RepositorySchedule
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		RepositoryStudent:  postgresql.NewStudentPostgres(db),
		RepositoryTeacher:  postgresql.NewTeacherPostgres(db),
		RepositoryAdmin:    postgresql.NewAdminPostgres(db),
		RepositorySchedule: postgresql.NewSchedulePostgres(db),
	}
}
