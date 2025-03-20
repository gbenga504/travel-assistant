package mongodb

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoDBCollection struct {
	collection *mongo.Collection
}

func NewMongoDBCollection(m *MongoDB, collection string) *MongoDBCollection {
	return &MongoDBCollection{
		collection: m.db.Collection(collection),
	}
}

func (c *MongoDBCollection) CreateOne() {}
