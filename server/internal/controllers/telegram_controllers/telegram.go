package telegram_controllers

import (
	"encoding/json"
	"fmt"
	"lit-log/internal/models/books"

	bolt "go.etcd.io/bbolt"
	tele "gopkg.in/telebot.v4"
)

func (h handler) StartUser(c tele.Context) error {

	userId := c.Sender().ID

	_ = h.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userIdJson, err := json.Marshal(userId)
		if err != nil {
			return err
		}

		user := bucket.Get(userIdJson)
		if user == nil {
			booksData := books.UserData{}

			booksDataJson, err := json.Marshal(booksData)
			if err != nil {
				return err
			}
			return bucket.Put(userIdJson, booksDataJson)
		}

		return nil
	})

	responseMessage := fmt.Sprintf("Welcome, %s", c.Sender().Username)

	return c.Send(responseMessage)
}
