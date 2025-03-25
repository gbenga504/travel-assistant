package threadrepository

import (
	"encoding/json"
	"time"

	"github.com/gbenga504/travel-assistant/utils/db"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
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
	Id        string `json:"id" bson:"_id"`
	UserId    string `json:"userId" bson:"userId"`
	ThreadId  string `json:"threadId" bson:"threadId"`
	GroupId   string `json:"groupId" bson:"groupId"`
	Role      `json:"role" bson:"role"`
	Content   []ThreadEntrySchemaContent `json:"content" bson:"content"`
	UpdatedAt time.Time                  `json:"updatedAt" bson:"updatedAt"`
	CreatedAt time.Time                  `json:"createdAt" bson:"createdAt"`
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

	if thread == nil {
		return []ThreadEntrySchema{}
	}

	return thread
}

type ThreadGroupSchema struct {
	Id      string              `json:"_id"`
	Entries []ThreadEntrySchema `json:"entries"`
}

func (r *ThreadRepository) GetThreadByIdWithGroupedEntries(threadId string) []ThreadGroupSchema {
	var result []ThreadGroupSchema

	aggData := r.collection.Aggregate(
		[]bson.D{
			{{Key: "$match", Value: bson.D{{Key: "threadId", Value: threadId}}}},
			{{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$groupId"},
				{Key: "entries", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
			}}},
		},
	)

	aggDataJSON, err := json.Marshal(aggData)

	if err != nil {
		logger.Fatal("Error when marschalling aggregated data in GetThreadByIdWithGroupedEntries", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrJSONParseIssue),
			Message:       errors.Message(errors.ErrJSONParseIssue),
			OriginalError: err.Error(),
		})
	}

	json.Unmarshal(aggDataJSON, &result)

	return result
}
