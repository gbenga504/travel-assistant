package lib

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gbenga504/travel-assistant/lib/routes"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Cannot load .env files")
	}
}

func startServer() {
	httpHandler := gin.New()
	routes.Routes(httpHandler)

	PORT, ok := os.LookupEnv("PORT")

	if !ok {
		logger.Error("Failed to load PORT", logger.ErrorOpt{
			Name:    errors.Name(errors.EnvLoadError),
			Message: errors.Message(errors.EnvLoadError),
		})
	}

	PORT = fmt.Sprintf(":%s", PORT)
	logger.Info(fmt.Sprintf("Listening on Port %s", PORT))

	panic(http.ListenAndServe(PORT, httpHandler))
}

func RunApp() {
	loadEnv()
	startServer()
}
