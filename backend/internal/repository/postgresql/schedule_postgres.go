package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type SchedulePostgres struct {
	db *gorm.DB
}

func (p *SchedulePostgres) Create(schedule []models.Schedule) error {
	return p.db.Create(schedule).Error
}

func (p *SchedulePostgres) Get() ([]models.Schedule, error) {
	var schedule []models.Schedule
	result := p.db.Find(&schedule)
	return schedule, result.Error
}

func (p *SchedulePostgres) Update(student []models.Schedule) error {
	return p.db.Updates(student).Error
}

func NewSchedulePostgres(db *gorm.DB) *SchedulePostgres {
	return &SchedulePostgres{db: db}
}
