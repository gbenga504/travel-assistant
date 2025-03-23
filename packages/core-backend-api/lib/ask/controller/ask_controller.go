package askcontroller

import (
	"io"
	"net/http"

	askservice "github.com/gbenga504/travel-assistant/lib/ask/service"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gin-gonic/gin"
)

type AskController struct {
	service *askservice.AskService
}

func NewAskController(service *askservice.AskService) *AskController {
	return &AskController{
		service,
	}
}

func (c *AskController) Post(ctx *gin.Context) {
	var reqBody struct {
		Query    string `json:"query" binding:"required"`
		ThreadId string `json:"threadId" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.Header("Content-Type", "application/json")

		ctx.JSON(
			http.StatusBadRequest,
			errors.ToErrorResponse(http.StatusText(http.StatusBadRequest), err.Error()),
		)

		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	output := make(chan string)
	done := make(chan bool)

	c.service.RunStream(
		reqBody.ThreadId,
		reqBody.Query,
		output,
		done,
	)

	ctx.Stream(func(w io.Writer) bool {
		select {
		case o := <-output:
			ctx.SSEvent("message", gin.H{
				"message": o,
			})
			return true
		case <-done:
			ctx.SSEvent("end_stream", gin.H{})
			return false
		case <-ctx.Writer.CloseNotify():
			return false
		}
	})
}
