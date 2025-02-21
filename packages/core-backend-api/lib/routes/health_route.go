package routes

import (
	"github.com/gbenga504/travel-assistant/lib/controllers"
	"github.com/gin-gonic/gin"
)

func healthRoute(httpHandler *gin.RouterGroup) {
	httpHandler.GET("/health", controllers.GetHealth)
}
