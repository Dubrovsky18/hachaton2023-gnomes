package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)


type StudentPostgres struct {
	db *gorm.DB
}

func (p *StudentPostgres) GetLogin(login string) (models.Student, error) {
	result := p.db.Where("login")
}

func (p *StudentPostgres) Create(student models.Student) error {
	return p.db.Create
}

func (p *StudentPostgres) Get(uuid int) (models.Student, error) {
	return p.db.Get(uuid)
}

func (p *StudentPostgres) Update(student models.Student) error {
	return p.db.
}

func NewStudentPostgres(db *gorm.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}
