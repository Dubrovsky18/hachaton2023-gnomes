package repository

import (
	"fmt"

	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/config"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(cfg *config.Config) (*postgresRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Admin{},
		&models.Audience{},
		&models.Student{},
		&models.Teacher{},
		&models.Group{},
		&models.Subject{},
		&models.Lesson{},
		&models.Schedule{},
	)
	if err != nil {
		panic(err)
	}

	pgSvc := &postgresRepository{db: db}

	// err = pgSvc.seeds()
	// if err != nil {
	// 	panic(err)
	// }

	return pgSvc, nil
}

// func (p *postgresRepository) seeds() error {
// 	cfg := config.GetConfig()

// 	var user models.User
// 	result := p.db.Where(models.User{Login: "admin"}).First(&user)
// 	if result.Error == nil || user.Id != 0 {
// 		return nil
// 	}

// 	seedUser := &models.User{
		
// 	}
// 	err = p.db.Create(seedUser)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (pg *postgresRepository) Create(table string, columns_data map[string]interface{}) error{
	return nil
}

func (pg *postgresRepository) Update(table string, id int, columns_data map[string]interface{}) error{
	return nil
}

func (pg *postgresRepository) Delete(table string, id int) error{
	return nil
}

func (pg *postgresRepository) GetSchedule(columns_args map[string]interface{}) ([]models.Schedule, error){
	return nil, nil
}

func (pg *postgresRepository) GetUser(table string, id int, email string, phone string) (*models.User, error){
	return nil, nil
}