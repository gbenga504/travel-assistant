package threadservice

import (
	"context"

	threadrepository "github.com/gbenga504/travel-assistant/lib/thread/repository"
	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	travelagent "github.com/gbenga504/travel-assistant/utils/travel_agent"
)

type ThreadService struct {
	respository  *threadrepository.ThreadRepository
	geminiClient *gemini.GeminiClient
}

func NewThreadService(repository *threadrepository.ThreadRepository, geminiClient *gemini.GeminiClient) *ThreadService {
	return &ThreadService{
		repository,
		geminiClient,
	}
}

func (s *ThreadService) RunStream(threadId string, query string, writer chan<- string, done chan<- bool) {
	ta := travelagent.SetupTravelAgent(s.geminiClient)
	thread := s.respository.GetThreadById(threadId)

	ta.History = convertThreadToHistories(thread)
	ta.ListenAndNotifyHistoryChange = func(h agent.History) {
		threadEntrySchema := convertHistoryToThreadEntrySchema(threadId, h)

		s.respository.CreateThreadEntry(&threadEntrySchema)
	}

	go func() {
		ta.RunStream(context.Background(), query, func(ctx context.Context, chunks []byte) {
			writer <- string(chunks)
		})

		done <- true
	}()
}
