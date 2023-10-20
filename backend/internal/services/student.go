package services

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
)

type StudentService struct {
	repositoryStudent repository.RepositoryStudent
}

func NewStudentService(student repository.RepositoryStudent) *StudentService {
	return &StudentService{
		repositoryStudent: student,
	}
}

func (p StudentService) Get(uuid int) (models.Student, error) {
	return p.repositoryStudent.Get(uuid)
}

func (p StudentService) GetLogin(login string) (models.Student, error) {
	return p.repositoryStudent.GetLogin(login)
}

func (p StudentService) Create(student models.Student) error {
	return p.repositoryStudent.Create(student)
}

func (p StudentService) Update(student models.Student) error {
	return p.repositoryStudent.Update(student)
}
