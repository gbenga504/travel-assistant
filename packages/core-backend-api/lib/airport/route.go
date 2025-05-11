package airport

import (
	airportcontroller "github.com/gbenga504/travel-assistant/lib/airport/controller"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes(httpHandler *gin.RouterGroup, controller *airportcontroller.AirportController) {
	httpHandler = httpHandler.Group("/airports")

	httpHandler.GET("/search", controller.Search)
}
