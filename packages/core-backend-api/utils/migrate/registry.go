package main

import (
	create_airports_1_2025_05_16T221702 "github.com/gbenga504/travel-assistant/utils/migrate/migrations/2025_05_16T221702_create_airports_1"
	create_airports_2_2025_05_17T142207 "github.com/gbenga504/travel-assistant/utils/migrate/migrations/2025_05_17T142207_create_airports_2"
	create_airports_3_2025_05_17T162703 "github.com/gbenga504/travel-assistant/utils/migrate/migrations/2025_05_17T162703_create_airports_3"
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
	addMigrationToRegistry("2025_05_17T142207_create_airports_2", &create_airports_2_2025_05_17T142207.CreateAirports220250517T142207{})
	addMigrationToRegistry("2025_05_17T162703_create_airports_3", &create_airports_3_2025_05_17T162703.CreateAirports320250517T162703{})
	//#registerMigration
}
