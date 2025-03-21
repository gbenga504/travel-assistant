package askservice

import (
	"context"

	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	travelagent "github.com/gbenga504/travel-assistant/utils/travel_agent"
)

type AskService struct {
	geminiClient *gemini.GeminiClient
}

func NewAskService(geminiClient *gemini.GeminiClient) *AskService {
	return &AskService{
		geminiClient,
	}
}

func (s *AskService) RunStream(query string, writer chan<- string, done chan<- bool) {
	agent := travelagent.SetupTravelAgent(s.geminiClient)

	go func() {
		agent.RunStream(context.Background(), query, func(ctx context.Context, chunks []byte) {
			writer <- string(chunks)
		})

		done <- true
	}()
}
