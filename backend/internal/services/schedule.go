package services

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/repository"
)

type ScheduleService struct {
	repositorySchedule repository.RepositorySchedule
}

func NewScheduleService(schedule repository.RepositorySchedule) *ScheduleService {
	return &ScheduleService{
		repositorySchedule: schedule,
	}
}

func (p ScheduleService) Get() ([]models.Schedule, error) {
	return p.repositorySchedule.Get(), nil
}

func (p ScheduleService) Create(schedule []models.Schedule) error {
	return p.repositorySchedule.Create(schedule)
}

func (p ScheduleService) Update(schedule []models.Schedule) error {
	return p.repositorySchedule.Update(schedule)
}
