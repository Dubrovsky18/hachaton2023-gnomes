package initializers

import "github.com/Dubrovsky18/hachaton2023-gnomes/internal/app/config"

func InitializeAppConfig() *config.AppConfig {
	return config.NewAppConfig()
}
