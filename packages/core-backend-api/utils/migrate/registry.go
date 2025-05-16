package main

import (
	create_airports_1_2025_05_16T221702 "github.com/gbenga504/travel-assistant/utils/migrate/migrations/2025_05_16T221702_create_airports_1"
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

	addMigrationToRegistry("2025_05_16T221702_create_airports_1", &create_airports_1_2025_05_16T221702.CreateAirports120250516T221702{})
	//#registerMigration
}
