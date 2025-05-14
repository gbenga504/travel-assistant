package main

import (
	"github.com/gbenga504/travel-assistant/utils/migrate/migrations"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

	addMigrationToRegistry("2025_05_13T205519_hello_world", &migrations.HelloWorld{})
	//#registerMigration
}
