
package migrations

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type HelloWorld struct{}

func (m *HelloWorld) Up(db *mongo.Database) error {
	// TODO: Implement migration up logic
	return nil
}

func (m *HelloWorld) Down(db *mongo.Database) error {
	// TODO: Implement migration down logic
	return nil
}
