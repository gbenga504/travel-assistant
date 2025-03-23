package mongodb

import (
	"context"
	"log"
	"reflect"

	"github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func (co *MongoDBCollection) CreateOne(document interface{}) {
	bsonD, documentRef := convertToBsonD(document)
	result, err := co.collection.InsertOne(
		context.Background(),
		bsonD,
	)

	if err != nil {
		logger.Fatal("CreateOne DB operation failed", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}

	if r, ok := result.InsertedID.(bson.ObjectID); ok {
		prop := documentRef.FieldByName("Id")
		prop.Set(reflect.ValueOf(r.Hex()))
	}
}

func (co *MongoDBCollection) FindMany(filter interface{}, documents interface{}) {
	// TODO: Find a generic way to handle filters to repository don't need to
	// know about its implementation detail
	cursor, err := co.collection.Find(context.Background(), filter)

	if err != nil {
		logger.Fatal("FindMany DB operation failed", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}

	err = cursor.All(context.Background(), documents)

	if err != nil {
		logger.Fatal("FindMany decoding DB operation failed", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}
}

func convertToBsonD(document interface{}) (result bson.D, documentRef reflect.Value) {
	ref := reflect.ValueOf(document)

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	// should double check we now have a struct (could still be anything)
	if ref.Kind() != reflect.Struct {
		log.Fatal("unexpected type")
	}

	for i := 0; i < ref.NumField(); i++ {
		// We don't want to include the Id field in the document
		if ref.Type().Field(i).Name == "Id" {
			continue
		}

		result = append(result, bson.E{
			Key:   utils.FirstLetterToLower(ref.Type().Field(i).Name),
			Value: ref.Field(i).Interface(),
		})
	}

	return result, ref
}
