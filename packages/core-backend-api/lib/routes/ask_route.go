package routes

import (
	"github.com/gbenga504/travel-assistant/lib/controllers"
	"github.com/gin-gonic/gin"
)

func askRoute(httpHandler *gin.RouterGroup) {
	httpHandler.POST("/ask", controllers.PostAsk)
}
