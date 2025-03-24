package thread

import (
	threadcontroller "github.com/gbenga504/travel-assistant/lib/thread/controller"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes(httpHandler *gin.RouterGroup, controller *threadcontroller.ThreadController) {
	httpHandler = httpHandler.Group("/thread")

	httpHandler.POST("/ask", controller.Post)
}
