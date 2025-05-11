package db

import "go.mongodb.org/mongo-driver/v2/bson"

type Db interface {
	Close()
	Collection(collection string) Collection
}

type FindManyOptions struct {
	Sort       *bson.D
	Limit      *int64
	Projection *bson.D
}

type Collection interface {
	CreateOneIndex(docment interface{}) bool
	CreateOne(document interface{})
	FindMany(filter interface{}, documents interface{}, findManyOpts *FindManyOptions)
	// TODO: Refactor technical debt. We should be DB agnostic
	Aggregate(aggregationFilter []bson.D) []bson.M
}
