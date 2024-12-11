package telegram

import (
	bolt "go.etcd.io/bbolt"
	tele "gopkg.in/telebot.v4"
)

//TODO

func (h handler) RegisterUser(c tele.Context) error {

	userId := c.Sender().ID

	_ = h.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		user := b.Get([]byte(string(rune(userId))))

		if user == nil {
			return nil
		}

		return nil
	})

	return c.Send("HE")
}
