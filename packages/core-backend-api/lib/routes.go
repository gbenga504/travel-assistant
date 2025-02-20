package lib

import (
	"github.com/gbenga504/travel-assistant/lib/health"
	"github.com/gin-gonic/gin"
)

func routes(httpHandler *gin.Engine) {
	v1 := httpHandler.Group("/api/v1")

	health.Routes(v1)
}
