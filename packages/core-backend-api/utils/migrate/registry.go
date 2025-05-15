package main

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	//#importMigration
)

type Migration interface {
	Up(db *mongo.Database) error
	Down(db *mongo.Database) error
}

var registry = make(map[string]Migration)

func addMigrationToRegistry(migrationName string, migration Migration) {
	registry[migrationName] = migration
}

func registerMigrations() {
	// Add migration registrations here

	//#registerMigration
}
