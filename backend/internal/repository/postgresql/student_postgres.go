package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type StudentPostgres struct {
	Db *gorm.DB
}

func (p *StudentPostgres) GetLogin(login string) (models.Student, error) {
	var student models.Student
	result := p.Db.Where("Email = ?", login).First(&student)
	return student, result.Error
}

func (p *StudentPostgres) Create(student models.Student) error {
	return p.Db.Create(student).Error
}

func (p *StudentPostgres) Get(uuid int) (models.Student, error) {
	var student models.Student
	result := p.Db.Where("id = ?", uuid).First(&student)
	return student, result.Error
}

func (p *StudentPostgres) Update(student models.Student) error {
	return p.Db.Updates(student).Error
}

func NewStudentPostgres(db *gorm.DB) *StudentPostgres {
	return &StudentPostgres{Db: db}
}
