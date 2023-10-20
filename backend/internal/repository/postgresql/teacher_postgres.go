package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type TeacherPostgres struct {
	db *gorm.DB
}

func (p *TeacherPostgres) GetLogin(login string) (models.Teacher, error) {
	var teacher models.Teacher
	result := p.db.Where("Email = ?", login).First(&teacher)
	return teacher, result.Error
}

func (p *TeacherPostgres) Create(teacher models.Teacher) error {
	return p.db.Create(teacher).Error
}

func (p *TeacherPostgres) Get(uuid int) (models.Teacher, error) {
	var teacher models.Teacher
	result := p.db.Where("id = ?", uuid).First(&teacher)
	return teacher, result.Error
}

func (p *TeacherPostgres) Update(teacher models.Teacher) error {
	result := p.db.Updates(teacher)
	return result.Error
}

func NewTeacherPostgres(db *gorm.DB) *TeacherPostgres {
	return &TeacherPostgres{db: db}
}
