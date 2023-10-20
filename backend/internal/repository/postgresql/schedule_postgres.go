package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type SchedulePostgres struct {
	db *gorm.DB
}

func (p *SchedulePostgres) Create(student models.Schedule) error {
	//TODO implement me
	panic("implement me")
}

func (p *SchedulePostgres) Get() ([]models.Schedule, error) {
	return p.db.Geta, nil
}

func (p *SchedulePostgres) Update(student models.Schedule) error {
	//TODO implement me
	panic("implement me")
}

func NewSchedulePostgres(db *gorm.DB) *SchedulePostgres {
	return &SchedulePostgres{db: db}
}
