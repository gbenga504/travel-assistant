package askservice

import (
	"context"

	askrepository "github.com/gbenga504/travel-assistant/lib/ask/repository"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	travelagent "github.com/gbenga504/travel-assistant/utils/travel_agent"
)

type AskService struct {
	respository  *askrepository.AskRepository
	geminiClient *gemini.GeminiClient
}

func NewAskService(repository *askrepository.AskRepository, geminiClient *gemini.GeminiClient) *AskService {
	return &AskService{
		repository,
		geminiClient,
	}
}

func (s *AskService) RunStream(query string, writer chan<- string, done chan<- bool) {
	s.respository.CreateChat()

	agent := travelagent.SetupTravelAgent(s.geminiClient)

	go func() {
		agent.RunStream(context.Background(), query, func(ctx context.Context, chunks []byte) {
			writer <- string(chunks)
		})

		done <- true
	}()
}
