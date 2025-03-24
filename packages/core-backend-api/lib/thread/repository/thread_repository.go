package threadrepository

import (
	"github.com/gbenga504/travel-assistant/utils/db"
	"go.mongodb.org/mongo-driver/v2/bson"
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

type ThreadEntrySchemaContent struct {
	Action  `json:"action" bson:"action"`
	Content string `json:"content" bson:"content"` // This should be a marschalled json of the action's content
}

type ThreadEntrySchema struct {
	Id       string `json:"id" bson:"_id"`
	UserId   string `json:"userId" bson:"userId"`
	ThreadId string `json:"threadId" bson:"threadId"`
	Role     `json:"role" bson:"role"`
	Content  []ThreadEntrySchemaContent `json:"content" bson:"content"`
}

type ThreadRepository struct {
	collection db.Collection
}

func NewThreadRepository(db db.Db) *ThreadRepository {
	return &ThreadRepository{
		collection: db.Collection("threads"),
	}
}

func (r *ThreadRepository) CreateThreadEntry(threadEntrySchema *ThreadEntrySchema) {
	r.collection.CreateOne(threadEntrySchema)
}

func (r *ThreadRepository) GetThreadById(id string) []ThreadEntrySchema {
	var thread []ThreadEntrySchema

	r.collection.FindMany(
		bson.D{{
			Key: "threadId",
			Value: bson.D{{
				Key:   "$eq",
				Value: id,
			}},
		}},
		&thread,
	)

	return thread
}
