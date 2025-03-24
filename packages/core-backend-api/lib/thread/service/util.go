package threadservice

import (
	threadrepository "github.com/gbenga504/travel-assistant/lib/thread/repository"
	"github.com/gbenga504/travel-assistant/utils/agent"
)

func convertHistoryToThreadEntrySchema(threadId string, history agent.History) threadrepository.ThreadEntrySchema {
	var threadEntrySchemaContent []threadrepository.ThreadEntrySchemaContent

	for _, c := range history.Content {
		threadEntrySchemaContent = append(threadEntrySchemaContent, threadrepository.ThreadEntrySchemaContent{
			Action:  threadrepository.Action(c.Action),
			Content: c.Content,
		})
	}

	return threadrepository.ThreadEntrySchema{
		ThreadId: threadId,
		Role:     threadrepository.Role(history.Role),
		Content:  threadEntrySchemaContent,
	}
}

func convertThreadToHistories(thread []threadrepository.ThreadEntrySchema) []*agent.History {
	var histories []*agent.History

	for _, th := range thread {
		var content []agent.HistoryContent

		for _, c := range th.Content {
			content = append(content, agent.HistoryContent{
				Action:  agent.Action(c.Action),
				Content: c.Content,
			})
		}

		histories = append(histories, &agent.History{
			Role:    agent.Role(th.Role),
			Content: content,
		})
	}

	return histories
}
