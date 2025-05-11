package mongodb

import (
	"context"
	"log"
	"reflect"
	"slices"
	"time"

	"github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/db"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDBCollection struct {
	collection *mongo.Collection
}

func NewMongoDBCollection(m *MongoDB, collection string) *MongoDBCollection {
	return &MongoDBCollection{
		collection: m.db.Collection(collection),
	}
}

func (co *MongoDBCollection) CreateOneIndex(document interface{}) bool {
	bsonD, _ := convertToBsonD(document, []string{})
	model := mongo.IndexModel{
		Keys: bsonD,
	}

	_, err := co.collection.Indexes().CreateOne(context.Background(), model)

	if err != nil {
		logger.Fatal("CreateOneIndex  DB operation failed", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}

	return true
}

func (co *MongoDBCollection) CreateOne(document interface{}) {
	bsonD, documentRef := convertToBsonD(document, []string{"Id", "CreatedAt", "UpdatedAt"})

	// Add UpdatedAt and CreatedAt to the bsonD
	bsonD = append(bsonD, bson.E{
		Key:   "createdAt",
		Value: time.Now(),
	}, bson.E{
		Key:   "updatedAt",
		Value: time.Now(),
	})

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

func (co *MongoDBCollection) FindMany(filter interface{}, documents interface{}, findManyOpts *db.FindManyOptions) {
	// TODO: Find a generic way to handle filters to repository don't need to
	// know about its implementation detail

	// Sort by ascending createdAt by default i.e oldest first to newest
	var opts *options.FindOptionsBuilder

	if findManyOpts == nil {
		opts = options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})
	} else {
		opts = options.Find()

		if findManyOpts.Sort != nil {
			opts = opts.SetSort(*findManyOpts.Sort)
		}

		if findManyOpts.Limit != nil {
			opts = opts.SetLimit(*findManyOpts.Limit)
		}

		if findManyOpts.Projection != nil {
			opts = opts.SetProjection(*findManyOpts.Projection)
		}
	}

	cursor, err := co.collection.Find(context.Background(), filter, opts)

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

func (co *MongoDBCollection) Aggregate(aggregationFilter []bson.D) []bson.M {
	cursor, err := co.collection.Aggregate(context.Background(), aggregationFilter)

	if err != nil {
		logger.Fatal("Aggregate DB operation failed", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}

	defer cursor.Close(context.Background())

	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		logger.Fatal("Aggregate DB operation failed: Cursor.All failed", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}

	return results
}

func convertToBsonD(document interface{}, ignoreFields []string) (result bson.D, documentRef reflect.Value) {
	ref := reflect.ValueOf(document)

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	// should double check we now have a struct (could still be anything)
	if ref.Kind() != reflect.Struct {
		log.Fatal("unexpected type. Wanted Struct")
	}

	for i := 0; i < ref.NumField(); i++ {
		// We don't want to include the fields that should be ignored in the bsonD
		if slices.Contains(ignoreFields, ref.Type().Field(i).Name) {
			continue
		}

		result = append(result, bson.E{
			Key:   utils.FirstLetterToLower(ref.Type().Field(i).Name),
			Value: ref.Field(i).Interface(),
		})
	}

	return result, ref
}
