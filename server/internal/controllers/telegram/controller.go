package telegram

import (
	bolt "go.etcd.io/bbolt"
	"gopkg.in/telebot.v4"
)

type handler struct {
	DB *bolt.DB
}

func RegisterRoutes(tg *telebot.Bot, db *bolt.DB) {
	h := &handler{
		DB: db,
	}

	tg.Handle("/start", h.RegisterUser)

}
