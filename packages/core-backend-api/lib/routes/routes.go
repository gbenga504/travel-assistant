package routes

import (
	"github.com/gbenga504/travel-assistant/lib/middlewares"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	"github.com/gin-gonic/gin"
)

func Routes(httpHandler *gin.Engine, geminiClient *gemini.GeminiClient) {
	// Apply global middlewares
	httpHandler.Use(middlewares.CORSMiddleware())

	v1 := httpHandler.Group("/api/v1")

	healthRoute(v1)
	askRoute(v1, geminiClient)
}
