package main

import (
	"fmt"
	"log/slog"

	"github.com/gbenga504/travel-assistant/lib"
	util "github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/joho/godotenv"
)

func setDefaultLogger() {
	slog.SetDefault(logger.NewLogger())
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		logger.Fatal("Cannot load .env files", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrEnvNotLoaded),
			Message:       errors.Message(errors.ErrEnvNotLoaded),
			OriginalError: err.Error(),
		})
	}
}

func main() {
	setDefaultLogger()
	loadEnv()

	PORT := fmt.Sprintf(":%s", util.LookupEnv("PORT"))
	logger.Info(fmt.Sprintf("Listening on Port %s", PORT))

	server := lib.NewServer(PORT)
	server.Run()
	server.ShutdownGracefully()
}
