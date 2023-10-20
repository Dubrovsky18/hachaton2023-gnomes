package repository

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/config"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/models"
)

type Repository interface {
	Create(table string, columns_data map[string]interface{}) error
	Update(table string, id int, columns_data map[string]interface{}) error
	Delete(table string, id int) error
	GetSchedule(columns_args map[string]interface{}) ([]models.Schedule, error)
	GetUser(table string, id int, email string, phone string) (*models.User, error)
}

func NewRepository(configuration *config.Config) Repository {
// 	switch configuration.App.DatabaseType {
// 	case "postgres":
	repo, err := NewPostgresRepository(configuration)
	if err != nil {
		panic(err)
	}
	return repo
	// }
	// return nil
}
