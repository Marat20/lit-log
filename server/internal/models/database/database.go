package database

import (
	"lit-log/internal/models/books"
	"lit-log/internal/pkg/config"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb(config config.Config, log *slog.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	err = db.AutoMigrate(&books.User{}, &books.Book{}, &books.ReadingProgress{})
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	
	log.Info("Connected to DB")
	return db
}
