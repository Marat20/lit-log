package telegram_controllers

import (
	tele "gopkg.in/telebot.v4"
)

//TODO

func (h handler) StartUser(c tele.Context) error {

	// userId := c.Sender().ID

	// _ = h.DB.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("Books"))
	// 	if b == nil {
	// 		return bolt.ErrBucketNotFound
	// 	}

	// 	userIdJson, err := json.Marshal(userId)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	user := b.Get(userIdJson)

	// 	booksData := []books.Book{}

	// 	booksDataJson, err := json.Marshal(booksData)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if user == nil {
	// 		return b.Put(userIdJson, booksDataJson)
	// 	}

	// 	return nil
	// })

	return c.Send("Welcome")
}
