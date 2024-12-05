package database

import (
	"lit-log/internal/pkg/config"
	"log/slog"

	bolt "go.etcd.io/bbolt"
)

func ConnectDB(cfg config.Config, log *slog.Logger) (*bolt.DB, error) {
	path := cfg.Db.Path
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	log.Info("Connected to database")

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Books"))
		if err != nil {
			log.Error(err.Error())
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte("Users"))
		if err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	})

	if err != nil {
		log.Error(err.Error())
	}

	return db, nil
}
