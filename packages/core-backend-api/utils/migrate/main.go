package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	loadEnv()
	registerMigrations()

	var (
		action        = flag.String("action", "", "Migration action (up|down)")
		migrationName = flag.String("migration", "", "Migration name (e.g., 2025-04-29T100421-add-new-users)")
	)
	flag.Parse()

	if *action == "" || *migrationName == "" {
		flag.Usage()
		os.Exit(1)
	}

	db, dbClient := connectDatabase()
	defer dbClient.Disconnect(context.Background())

	migration, exists := registry[*migrationName]
	if !exists {
		log.Fatalf("Migration '%s' not registered", *migrationName)
	}

	switch *action {
	case "up":
		runMigration(db, *migrationName, migration)

	case "down":
		rollbackMigration(db, *migrationName, migration)

	default:
		log.Fatal("Invalid action. Use 'up', 'down'")
	}

	fmt.Println("Operation completed successfully")
}

func recordMigration(db *mongo.Database, migrationName string) error {
	_, err := db.Collection("migrations").InsertOne(context.Background(), bson.D{
		{Key: "name", Value: migrationName},
	})

	return err
}

func removeMigration(db *mongo.Database, migrationName string) error {
	_, err := db.Collection("migrations").DeleteOne(context.Background(), bson.D{
		{Key: "name", Value: migrationName},
	})

	return err
}

func isApplied(db *mongo.Database, migrationName string) bool {
	res := db.Collection("migrations").FindOne(context.Background(), bson.D{
		{Key: "name", Value: migrationName},
	})

	return res.Err() == nil
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Cannot load .env files. Error: %s", err.Error())
	}
}

func connectDatabase() (db *mongo.Database, dbClient *mongo.Client) {
	uri, ok := os.LookupEnv("MONGODB_URI")

	if !ok {
		log.Fatal("MONGODB_URI does not exist in .env")
	}

	dbClient, err := mongo.Connect(options.Client().
		ApplyURI(uri).
		SetBSONOptions(&options.BSONOptions{
			// Automatically convert ObjectID to hex string
			ObjectIDAsHexString: true,
		}),
	)

	if err != nil {
		log.Fatalf("Cannot connect to MongoDB. Error: %s", err.Error())
	}

	dbName, ok := os.LookupEnv("DATABASE_NAME")

	if !ok {
		log.Fatal("DATABASE_NAME does not exist in .env")
	}

	db = dbClient.Database(dbName)

	return db, dbClient
}

func runMigration(db *mongo.Database, migrationName string, migration Migration) {
	if isApplied(db, migrationName) {
		log.Printf("Migration %s already applied", migrationName)

		return
	}

	fmt.Printf("Applying migration: %s\n", migrationName)

	if err := migration.Up(db); err != nil {
		log.Fatalf("Migration failed: %s", err.Error())
	}

	if err := recordMigration(db, migrationName); err != nil {
		log.Fatalf("Failed to record migration: %s", err.Error())
	}
}

func rollbackMigration(db *mongo.Database, migrationName string, migration Migration) {
	if !isApplied(db, migrationName) {
		log.Printf("Migration %s not applied", migrationName)

		return
	}

	fmt.Printf("Rolling back migration: %s\n", migrationName)

	if err := migration.Down(db); err != nil {
		log.Fatalf("Rollback failed: %s", err.Error())
	}

	if err := removeMigration(db, migrationName); err != nil {
		log.Fatalf("Failed to remove migration record: %s", err.Error())
	}
}
