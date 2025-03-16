package lib

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gbenga504/travel-assistant/lib/routes"
	util "github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
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
	GEMINI_API_KEY := util.LookupEnv("GEMINI_API_KEY")
	geminiClient := gemini.NewClient(context.Background(), GEMINI_API_KEY)
	defer geminiClient.Close()

	httpHandler := gin.New()
	routes.Routes(httpHandler, geminiClient)

	PORT := fmt.Sprintf(":%s", util.LookupEnv("PORT"))
	logger.Info(fmt.Sprintf("Listening on Port %s", PORT))

	panic(http.ListenAndServe(PORT, httpHandler))
}

func RunApp() {
	loadEnv()
	startServer()
}
