package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const migrationTemplate = `
package migrations

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func init() {
	RegisterMigration("%s", &%s{})
}

type %s struct{}

func (m *%s) Up(db *mongo.Database) error {
	// TODO: Implement migration up logic
	return nil
}

func (m *%s) Down(db *mongo.Database) error {
	// TODO: Implement migration down logic
	return nil
}
`

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run create_migration.go <migration_name>")
		os.Exit(1)
	}

	migrationName := strings.Join(os.Args[1:], "_")
	migrationName = strings.ToLower(strings.ReplaceAll(migrationName, " ", "_"))
	timestamp := time.Now().Format("2006_01_02T150405")

	filename := fmt.Sprintf("%s_%s", timestamp, migrationName)
	filePath := fmt.Sprintf("migrations/%s.go", filename)

	structName := toCamelCase(migrationName)
	content := fmt.Sprintf(migrationTemplate, filename, structName, structName, structName, structName)

	// Ensure migrations directory exists
	if _, err := os.Stat("migrations"); os.IsNotExist(err) {
		if err := os.Mkdir("migrations", 0755); err != nil {
			fmt.Println("Failed to create migrations directory:", err)
			os.Exit(1)
		}
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		fmt.Println("Failed to write migration file:", err)
		os.Exit(1)
	}

	fmt.Println("Created migration:", filename)
}
