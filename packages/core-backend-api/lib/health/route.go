package health

import (
	healthcontroller "github.com/gbenga504/travel-assistant/lib/health/controller"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes(httpHandler *gin.RouterGroup, controller *healthcontroller.HealthController) {
	httpHandler = httpHandler.Group("/health")

	httpHandler.GET("", controller.Get)
}
