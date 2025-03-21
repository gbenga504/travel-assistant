package ask

import (
	askcontroller "github.com/gbenga504/travel-assistant/lib/ask/controller"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes(httpHandler *gin.RouterGroup, controller *askcontroller.AskController) {
	httpHandler = httpHandler.Group("/ask")

	httpHandler.POST("", controller.Post)
}
