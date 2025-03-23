package mongodb

import (
	"context"

	util "github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/db"
	"github.com/gbenga504/travel-assistant/utils/errors"
	"github.com/gbenga504/travel-assistant/utils/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

var _ db.Db = (*MongoDB)(nil)

func Connect(dbName string) *MongoDB {
	uri := util.LookupEnv("MONGODB_URI")
	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		logger.Fatal("Cannot connect to MongoDB", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}

	db := client.Database(dbName)

	return &MongoDB{
		client,
		db,
	}
}

func (m *MongoDB) Close() {
	if err := m.client.Disconnect(context.Background()); err != nil {
		logger.Fatal("Cannot close MongoDB connection", logger.ErrorOpt{
			Name:          errors.Name(errors.ErrDatabaseIssue),
			Message:       errors.Message(errors.ErrDatabaseIssue),
			OriginalError: err.Error(),
		})
	}
}

func (m *MongoDB) Collection(collection string) db.Collection {
	return NewMongoDBCollection(m, collection)
}
