package lib

import (
	"fmt"
	"net/http"
	"os"

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
	httpHandler := gin.Default()
	routes(httpHandler)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	logger.Info(logger.LoggerParams{
		Msg: fmt.Sprintf("Listening on Port %s", PORT),
	})

	panic(http.ListenAndServe(PORT, httpHandler))
}

func RunApp() {
	loadEnv()
	startServer()
}
