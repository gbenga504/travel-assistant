package db

import "go.mongodb.org/mongo-driver/v2/bson"

type Db interface {
	Close()
	Collection(collection string) Collection
}

type Collection interface {
	CreateOne(document interface{})
	FindMany(filter interface{}, documents interface{})
	// TODO: Refactor technical debt. We should be DB agnostic
	Aggregate(aggregationFilter []bson.D) []bson.M
}
