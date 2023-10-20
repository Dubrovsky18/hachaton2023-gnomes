package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type StudentPostgres struct {
	db *gorm.DB
}

func (p *StudentPostgres) GetLogin(login string) (models.Student, error) {
	var student models.Student
	result := p.db.Where("Email = ?", login).First(&student)
	return student, result.Error
}

func (p *StudentPostgres) Create(student models.Student) error {
	return p.db.Create(student).Error
}

func (p *StudentPostgres) Get(uuid int) (models.Student, error) {
	var student models.Student
	result := p.db.Where("id = ?", uuid).First(&student)
	return student, result.Error
}

func (p *StudentPostgres) Update(student models.Student) error {
	return p.db.Updates(student).Error
}

func NewStudentPostgres(db *gorm.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}
