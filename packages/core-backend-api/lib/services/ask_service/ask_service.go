package askservice

import (
	"context"

	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	travelagent "github.com/gbenga504/travel-assistant/utils/travel_agent"
)

func RunStream(query string, gc *gemini.GeminiClient, writer chan<- string, done chan<- bool) {
	agent := travelagent.SetupTravelAgent(gc)

	go func() {
		agent.RunStream(context.Background(), query, func(ctx context.Context, chunks []byte) {
			writer <- string(chunks)
		})

		done <- true
	}()
}
