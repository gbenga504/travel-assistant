package health

import (
	"github.com/gin-gonic/gin"
)

func Routes(httpHandler *gin.RouterGroup) {
	httpHandler.GET("/health", Get)
}
