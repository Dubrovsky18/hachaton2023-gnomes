package repository

import (
	"fmt"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() *gorm.DB {

	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
	//
	//	//dsn := "postgres://postgres:postgres@95.163.211.20:5432/postgres?sslmode=disable"
	//	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//	if err != nil {
	//		log.Fatalf("failed to connect database" + err.Error())
	//	}
	//	err = db.AutoMigrate(
	//		&models.Teacher{},
	//		&models.Admin{},
	//		&models.Schedule{},
	//		&models.Lesson{},
	//		&models.Subject{},
	//		&models.Group{},
	//		&models.Audience{},
	//	)
	//	if err != nil {
	//		log.Fatalf("failed Migrate")
	//	}
	//
	//	return db
	//
	//}
}
