package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	handlers "github.com/ummuys/effective_mobile_ts/handlers/subscription"
	"github.com/ummuys/effective_mobile_ts/logger"
	"github.com/ummuys/effective_mobile_ts/repository"
	"github.com/ummuys/effective_mobile_ts/router"
	service "github.com/ummuys/effective_mobile_ts/service/subscription"
)

func main() {

	//---ENV
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	//---LOGGER
	logger, err := logger.InitLogger("logs/")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info().Msg("starting the programm...")

	logger.Info().Msg("env file is loaded")

	//---DATABASES
	db, err := repository.NewDatabase(logger)
	if err != nil {
		logger.Err(err)
		log.Fatal(err)
	}
	logger.Info().Msg("databases initialized")

	//---SERVICES
	followService := service.NewSubsService(db, logger)
	logger.Info().Msg("service initialized")

	//---HANDLERS
	followHandler := handlers.NewSubsHandler(followService, logger)
	logger.Info().Msg("handlers initialized")

	//---CHANEL FOR GRACEFUL SHOTDOWN
	chSD := make(chan os.Signal, 1)
	signal.Notify(chSD, syscall.SIGINT, syscall.SIGTERM)

	//---SERVER
	serv := router.CreateServer(followHandler)
	go router.RunServer(serv, logger)

	//---WAIT SHUTDOWN SIGNAL
	<-chSD
	logger.Info().Msg("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//---SHOTDOWN
	servErr := serv.Shutdown(ctx)
	if servErr != nil {
		msg := fmt.Sprintf("shotdown err: %v", servErr)
		logger.Fatal().Msg(msg)
	}

	dbErr := db.Close()
	if dbErr != nil {
		logger.Fatal().Msg(dbErr.Error())
	}

	if dbErr != nil || servErr != nil {
		logger.Fatal().Msg("server shutdown with errors")
	} else {
		logger.Info().Msg("server shutdown gracefully")
	}

}
