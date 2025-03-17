package routes

import (
	"github.com/gbenga504/travel-assistant/lib/controllers"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	"github.com/gin-gonic/gin"
)

func askRoute(httpHandler *gin.RouterGroup, gc *gemini.GeminiClient) {
	httpHandler.POST("/ask", func(ctx *gin.Context) {
		controllers.PostAsk(ctx, gc)
	})
}
