package main

import (
	"lit-log/internal/controllers"
	"lit-log/internal/pkg/database"
	"lit-log/internal/pkg/config"
	"lit-log/internal/pkg/logging"
	"lit-log/internal/services/server"
)

func main() {

	log, err := logging.New()
	if err != nil {
		log.Error(err.Error())
		return
	}

	defer log.Info("Shutting down...")

	config := config.LoadConfig()

	db := database.ConnectDB(*config, log)

	server := server.New()

	controllers.RegisterRoutes(server, db)

	server.Run()
}
