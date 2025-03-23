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
