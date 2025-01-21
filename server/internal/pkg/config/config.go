package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db      DbConfig
	TgToken string
}

type DbConfig struct {
	Path string
}

func LoadConfig() *Config {

	godotenv.Load()

	return &Config{
		Db: DbConfig{
			Path: os.Getenv("DB_PATH"),
		},
		TgToken: os.Getenv("TELEGRAM_API_TOKEN"),
	}
}
