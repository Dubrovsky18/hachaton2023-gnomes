package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type TeacherPostgres struct {
	db *gorm.DB
}

func (p *TeacherPostgres) GetLogin(login string) (models.Student, error) {
	//TODO implement me
	panic("implement me")
}

func (p *TeacherPostgres) Create(teacher models.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func (p *TeacherPostgres) Get(uuid int) (models.Teacher, error) {
	return p.db.Get(uuid)
}

func (p *TeacherPostgres) Update(teacher models.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func NewTeacherPostgres(db *gorm.DB) *TeacherPostgres {
	return &TeacherPostgres{db: db}
}
