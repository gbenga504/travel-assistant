package askrepository

import (
	"github.com/gbenga504/travel-assistant/utils/db"
)

type Role string
type Action string

const (
	// Roles detailing the type of users in the system
	AIRole     Role = "ai"
	UserRole   Role = "user"
	SystemRole Role = "system"
)

const (
	// Actions detailing the type of actions that can be performed by any of the roles
	ToolCallAction     Action = "toolCall"
	ToolResponseAction Action = "toolResponse"
	TextAction         Action = "text"
)

type ChatSchemaContent struct {
	Action
	Content string // This should be a marschalled json of the action's content
}

type ChatSchema struct {
	Id       string
	UserId   string
	ThreadId string
	Role
	Content []ChatSchemaContent
}

type AskRepository struct {
	collection db.Collection
}

func NewAskRepository(db db.Db) *AskRepository {
	return &AskRepository{
		collection: db.Collection("chat"),
	}
}

func (r *AskRepository) CreateChat(chatSchema *ChatSchema) {
	r.collection.CreateOne(chatSchema)
}
