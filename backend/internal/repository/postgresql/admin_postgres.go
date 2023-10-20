package postgresql

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/gorm"
)

type AdminPostgres struct {
	db *gorm.DB
}

func (p *AdminPostgres) GetLogin(login string) (models.Admin, error) {
	var admin models.Admin
	result := p.db.Where("Email = ?", login).First(&admin)
	return admin, result.Error
}

func (p *AdminPostgres) Create(admin models.Admin) error {
	return p.db.Create(admin).Error
}

func (p *AdminPostgres) Get(uuid int) (models.Admin, error) {
	var admin models.Admin
	result := p.db.Where("id = ?", uuid).First(&admin)
	return admin, result.Error
}

func (p *AdminPostgres) Update(admin models.Admin) error {
	result := p.db.Updates(admin)
	return result.Error
}

func NewAdminPostgres(db *gorm.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}
