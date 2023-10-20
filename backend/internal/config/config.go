package config

import (
	"github.com/gobuffalo/envy"
)

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	DBName   string
	SSLMode  string
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}

	return &Config{
		Database: Database{
			Host:     envy.Get("POSTGRES_HOST", "localhost"),
			Port:     envy.Get("POSTGRES_PORT", "5431"),
			User:     envy.Get("POSTGRES_USER", "postgres"),
			Password: envy.Get("POSTGRES_PASSWORD", "postgres"),
			DBName:   envy.Get("POSTGRES_NAME", "postgres"),
			SSLMode:  envy.Get("SSLMode", "disable"),
		},
	}
}
