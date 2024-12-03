package telegram

import (
	"lit-log/internal/pkg/config"
	"log/slog"
	"time"

	tele "gopkg.in/telebot.v4"
)

func New(config *config.Config, log *slog.Logger) *tele.Bot {
	pref := tele.Settings{
		Token:  config.TgToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	log.Info("Telegram bot started")

	return b
}
