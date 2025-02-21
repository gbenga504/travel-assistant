package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(httpHandler *gin.Engine) {
	v1 := httpHandler.Group("/api/v1")

	healthRoute(v1)
}
