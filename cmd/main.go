package main

import (
	"log"

	"github.com/joho/godotenv"
	handlers "github.com/ummuys/effective_mobile_ts/handlers/subscription"
	"github.com/ummuys/effective_mobile_ts/logger"
	"github.com/ummuys/effective_mobile_ts/repository"
	"github.com/ummuys/effective_mobile_ts/router"
	service "github.com/ummuys/effective_mobile_ts/service/subscription"
)

func main() {
	logger := logger.InitLogger("logs/")
	logger.Info().Msg("Start the programm . . .")

	err := godotenv.Load(".env")
	if err != nil {
		logger.Err(err)
		log.Fatal(err)
	}
	logger.Info().Msg("Env file is loaded")

	//---DATABASES
	db, err := repository.NewDatabase(logger)
	if err != nil {
		logger.Err(err)
		log.Fatal(err)
	}
	logger.Info().Msg("Db is setted")

	//---SERVICES
	followService := service.NewSubsService(db, logger)
	logger.Info().Msg("Service is setted")

	//---HANDLERS
	followHandler := handlers.NewSubsHandler(followService, logger)
	logger.Info().Msg("Handlers is setted")

	router.RunRouter(followHandler)
}
