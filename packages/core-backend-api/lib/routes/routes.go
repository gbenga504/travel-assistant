package routes

import (
	"github.com/gbenga504/travel-assistant/lib/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(httpHandler *gin.Engine) {
	// Apply global middlewares
	httpHandler.Use(middlewares.CORSMiddleware())

	v1 := httpHandler.Group("/api/v1")

	healthRoute(v1)
	askRoute(v1)
}
