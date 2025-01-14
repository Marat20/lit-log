package main

import (
	"flag"
	"fmt"
	bookControllers "lit-log/internal/controllers/book_controllers"
	telegramControllers "lit-log/internal/controllers/telegram_controllers"
	"lit-log/internal/pkg/config"
	"lit-log/internal/pkg/database"
	"lit-log/internal/pkg/logging"
	"lit-log/internal/services/server"
	"lit-log/internal/services/telegram"
)

// //go:embed asset
// var fs embed.FS

func main() {

	// x, _ := fs.ReadFile("asset/in")

	port := flag.Int("port", 8080, "serve port")
	flag.Parse()

	log, err := logging.New()
	if err != nil {
		log.Error(err.Error())
		return
	}

	defer log.Info("Shutting down...")

	config := config.LoadConfig()

	tg := telegram.New(config, log)

	go tg.Start()

	db, err := database.ConnectDB(*config, log)
	if err != nil {
		log.Error(err.Error())
		return
	}

	defer db.Close()

	server := server.New()

	bookControllers.RegisterRoutes(server, db)

	telegramControllers.RegisterRoutes(tg, db)

	server.Run(fmt.Sprintf(":%d", *port))
}
