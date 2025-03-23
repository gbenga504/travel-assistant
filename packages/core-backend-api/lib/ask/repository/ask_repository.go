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
	Action  `json:"action" bson:"action"`
	Content string `json:"content" bson:"content"` // This should be a marschalled json of the action's content
}

type ChatSchema struct {
	Id       string `json:"id" bson:"_id"`
	UserId   string `json:"userId" bson:"userId"`
	ThreadId string `json:"threadId" bson:"threadId"`
	Role     `json:"role" bson:"role"`
	Content  []ChatSchemaContent `json:"content" bson:"content"`
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
