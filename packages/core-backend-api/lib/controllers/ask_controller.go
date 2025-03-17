package controllers

import (
	"io"
	"net/http"

	askservice "github.com/gbenga504/travel-assistant/lib/services/ask_service"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gin-gonic/gin"
)

func PostAsk(c *gin.Context, gc *gemini.GeminiClient) {
	var reqBody struct {
		Query string `json:"query" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(
			http.StatusBadRequest,
			errors.ToErrorResponse(http.StatusText(http.StatusBadRequest), err.Error()),
		)
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	output := make(chan string)
	done := make(chan bool)

	askservice.RunStream(
		reqBody.Query,
		gc,
		output,
		done,
	)

	c.Stream(func(w io.Writer) bool {
		select {
		case o := <-output:
			c.SSEvent("message", gin.H{
				"message": o,
			})
			return true
		case <-done:
			c.SSEvent("end_stream", gin.H{})
			return false
		case <-c.Writer.CloseNotify():
			return false
		}
	})
}
