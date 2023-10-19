package services

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
)

type TemplateService interface {
}

type TemplateServiceImpl struct {
	repos    repository.TemplateRepository
	template *models.Template
}

func NewTemplateRepository(repos repository.TemplateRepository) *TemplateServiceImpl {
	return &TemplateServiceImpl{repos: repos}
}
