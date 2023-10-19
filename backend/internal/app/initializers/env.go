package initializers

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"
	"github.com/gobuffalo/envy"
)

func InitializeEnv() {
	if err := envy.Load(); err != nil {
		logger.Info("can not load .env file", "error", err)
		envy.Reload()
	}
}
