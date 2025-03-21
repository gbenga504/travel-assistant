package mongodb

import (
	"context"

	util "github.com/gbenga504/travel-assistant/utils"
	"github.com/gbenga504/travel-assistant/utils/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func Connect() *MongoDB {
	uri := util.LookupEnv("MONGODB_URI")
	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		// TODO panic
	}

	// The database will be passed in the connection string, hence we skip it here
	db := client.Database("")

	return &MongoDB{
		client,
		db,
	}
}

func (m *MongoDB) Close() {
	if err := m.client.Disconnect(context.Background()); err != nil {
		// TODO panic
	}
}

func (m *MongoDB) Collection(collection string) db.Collection {
	return NewMongoDBCollection(m, collection)
}
