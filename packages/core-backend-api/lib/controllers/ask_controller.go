package controllers

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	travelagent "github.com/gbenga504/travel-assistant/utils/travel_agent"
	"github.com/gin-gonic/gin"
)

func PostAsk(c *gin.Context, geminiClient *gemini.GeminiClient) {
	var reqBody struct {
		Query string `json:"query" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    map[string]any{"name": http.StatusText(http.StatusBadRequest), "message": err.Error()},
		})
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	ta := travelagent.SetupTravelAgent(geminiClient)
	llmOutput := make(chan string)
	quit := make(chan bool)

	go func() {
		ta.RunStream(c.Request.Context(), reqBody.Query, func(ctx context.Context, chunks []byte) {
			llmOutput <- string(chunks)
		})

		quit <- true
	}()

	c.Stream(func(w io.Writer) bool {
		select {
		case o := <-llmOutput:
			c.SSEvent("message", gin.H{
				"message": o,
			})
			return true
		case <-quit:
			c.SSEvent("EndStream", gin.H{})
			return false
		case <-c.Writer.CloseNotify():
			return false
		}
	})
}
