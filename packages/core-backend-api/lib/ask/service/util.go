package askservice

import (
	askrepository "github.com/gbenga504/travel-assistant/lib/ask/repository"
	"github.com/gbenga504/travel-assistant/utils/agent"
)

func convertHistoryToChatSchema(threadId string, history agent.History) askrepository.ChatSchema {
	var chatSchemaContent []askrepository.ChatSchemaContent

	for _, c := range history.Content {
		chatSchemaContent = append(chatSchemaContent, askrepository.ChatSchemaContent{
			Action:  askrepository.Action(c.Action),
			Content: c.Content,
		})
	}

	return askrepository.ChatSchema{
		ThreadId: threadId,
		Role:     askrepository.Role(history.Role),
		Content:  chatSchemaContent,
	}
}

func convertChatSchemasToHistories(chatSchemas []askrepository.ChatSchema) []*agent.History {
	var histories []*agent.History

	for _, cs := range chatSchemas {
		var content []agent.HistoryContent

		for _, c := range cs.Content {
			content = append(content, agent.HistoryContent{
				Action:  agent.Action(c.Action),
				Content: c.Content,
			})
		}

		histories = append(histories, &agent.History{
			Role:    agent.Role(cs.Role),
			Content: content,
		})
	}

	return histories
}
