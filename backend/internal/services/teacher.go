package services

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
)

type TeacherService struct {
	repositoryTeacher repository.RepositoryTeacher
}

func NewTeacherService(teacher repository.RepositoryTeacher) *TeacherService {
	return &TeacherService{
		repositoryTeacher: teacher,
	}
}

func (p TeacherService) GetLogin(login string) (models.Teacher, error) {
	return p.repositoryTeacher.GetLogin(login)
}

func (p TeacherService) Get(uuid int) (models.Teacher, error) {
	return p.repositoryTeacher.Get(uuid)
}

func (p TeacherService) Create(teacher models.Teacher) error {
	return p.repositoryTeacher.Create(teacher)
}

func (p TeacherService) Update(teacher models.Teacher) error {
	return p.repositoryTeacher.Update(teacher)
}
