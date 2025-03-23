package askservice

import (
	"context"

	askrepository "github.com/gbenga504/travel-assistant/lib/ask/repository"
	"github.com/gbenga504/travel-assistant/utils/agent"
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

func (s *AskService) RunStream(threadId string, query string, writer chan<- string, done chan<- bool) {
	ta := travelagent.SetupTravelAgent(s.geminiClient)

	ta.ListenAndNotifyHistoryChange = func(h agent.History) {
		chatSchema := convertHistoryToChatSchema(threadId, h)

		s.respository.CreateChat(&chatSchema)
	}

	go func() {
		ta.RunStream(context.Background(), query, func(ctx context.Context, chunks []byte) {
			writer <- string(chunks)
		})

		done <- true
	}()
}
