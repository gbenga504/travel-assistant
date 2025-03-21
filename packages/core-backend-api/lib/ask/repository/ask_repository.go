package askrepository

import (
	"github.com/gbenga504/travel-assistant/utils/db"
)

type AskRepository struct {
	collection db.Collection
}

func NewAskRepository(db db.Db) *AskRepository {
	return &AskRepository{
		collection: db.Collection("chat"),
	}
}

func (r *AskRepository) CreateChat() {
	r.collection.CreateOne()
}
