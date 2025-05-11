package airportcontroller

import (
	"net/http"

	airportservice "github.com/gbenga504/travel-assistant/lib/airport/service"
	"github.com/gbenga504/travel-assistant/utils"
	"github.com/gin-gonic/gin"
)

type AirportController struct {
	service *airportservice.AirportService
}

func NewAirportController(service *airportservice.AirportService) *AirportController {
	return &AirportController{
		service,
	}
}

// Search is a special controller because it relies on a complex
// pattern
func (c *AirportController) Search(ctx *gin.Context) {
	searchTerm := ctx.Query("q")
	result := c.service.SearchAirports(searchTerm)

	ctx.JSON(http.StatusOK, utils.ToSuccessResponse(result))
}
