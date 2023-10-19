package repository

import (
	"fmt"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/config"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TemplateRepository interface {
}

type TemplateRepositoryImpl struct {
	db *gorm.DB
}

func NewTemplateRepository() TemplateRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	pgSvc := &TemplateRepositoryImpl{db: db}
	err = db.AutoMigrate(&models.Template{})
	if err != nil {
		panic(err)
	}
	return pgSvc
}
