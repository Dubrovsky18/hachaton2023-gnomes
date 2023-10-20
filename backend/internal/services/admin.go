package services

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
)

type AdminService struct {
	repositoryAdmin repository.RepositoryAdmin
}

func NewAdminService(admin repository.RepositoryAdmin) *AdminService {
	return &AdminService{
		repositoryAdmin: admin,
	}
}

func (p AdminService) Get(uuid int) (models.Admin, error) {
	return p.repositoryAdmin.Get(uuid)
}

func (p AdminService) GetLogin(login string) (models.Admin, error) {
	return p.repositoryAdmin.GetLogin(login)
}

func (p AdminService) Create(student models.Admin) error {
	return p.repositoryAdmin.Create(student)
}

func (p AdminService) Update(student models.Admin) error {
	return p.repositoryAdmin.Update(student)
}
