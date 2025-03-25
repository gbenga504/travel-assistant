package threadservice

import (
	"context"

	threadrepository "github.com/gbenga504/travel-assistant/lib/thread/repository"
	"github.com/gbenga504/travel-assistant/utils/agent"
	"github.com/gbenga504/travel-assistant/utils/agent/llms/gemini"
	travelagent "github.com/gbenga504/travel-assistant/utils/travel_agent"
	gonanoid "github.com/matoous/go-nanoid/v2"
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
	nanoid, _ := gonanoid.New()

	ta := travelagent.SetupTravelAgent(s.geminiClient)
	thread := s.respository.GetThreadById(threadId)

	ta.History = convertThreadToHistories(thread)
	ta.ListenAndNotifyHistoryChange = func(h agent.History) {
		threadEntrySchema := convertHistoryToThreadEntrySchema(threadId, nanoid, h)

		s.respository.CreateThreadEntry(&threadEntrySchema)
	}

	go func() {
		ta.RunStream(context.Background(), query, func(ctx context.Context, chunks []byte) {
			writer <- string(chunks)
		})

		done <- true
	}()
}

type GroupedThreadEntry struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func (s *ThreadService) GetThreadByIdWithGroupedEntries(id string) []GroupedThreadEntry {
	var result []GroupedThreadEntry

	for _, schema := range s.respository.GetThreadByIdWithGroupedEntries(id) {
		groupedThreadEntry := &GroupedThreadEntry{
			Id: schema.Id,
		}

		// Find the question in the entries
		for _, entry := range schema.Entries {
			switch entry.Role {
			case threadrepository.UserRole:
				groupedThreadEntry.Question = entry.Content[0].Content

			case threadrepository.AIRole:
				groupedThreadEntry.Answer = entry.Content[0].Content
			}
		}

		result = append(result, *groupedThreadEntry)
	}

	if len(result) == 0 {
		return []GroupedThreadEntry{}
	}

	return result
}
