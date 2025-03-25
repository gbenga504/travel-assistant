package threadcontroller

import (
	"io"
	"net/http"

	threadservice "github.com/gbenga504/travel-assistant/lib/thread/service"
	"github.com/gbenga504/travel-assistant/utils"
	"github.com/gin-gonic/gin"
)

type ThreadController struct {
	service *threadservice.ThreadService
}

func NewThreadController(service *threadservice.ThreadService) *ThreadController {
	return &ThreadController{
		service,
	}
}

func (c *ThreadController) Post(ctx *gin.Context) {
	var reqBody struct {
		Query    string `json:"query" binding:"required"`
		ThreadId string `json:"threadId" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.Header("Content-Type", "application/json")

		ctx.JSON(
			http.StatusBadRequest,
			utils.ToErrorResponse(http.StatusText(http.StatusBadRequest), err.Error()),
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

func (c *ThreadController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	result := c.service.GetThreadByIdWithGroupedEntries(id)

	ctx.JSON(http.StatusOK, utils.ToSuccessResponse(result))
}
