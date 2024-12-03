package database

import (
	"lit-log/internal/models/books"
	"lit-log/internal/pkg/config"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config, log *slog.Logger) *gorm.DB {
	dsn := cfg.Db.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("error connecting to database: " + err.Error())
	}

	if err := db.AutoMigrate(&books.Book{}); err != nil {
		log.Error("error migrating database schema: " + err.Error())
	}

	log.Info("Connected to DB")

	return db
}
