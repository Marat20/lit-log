package logging

import (
	"log/slog"
	"os"
)

func New() (*slog.Logger, error) {
	handlerOptions := slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, &handlerOptions)
	logging := slog.New(handler)

	return logging, nil
}
