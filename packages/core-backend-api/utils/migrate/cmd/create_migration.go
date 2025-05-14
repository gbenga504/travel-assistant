package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const migrationTemplate = `
package migrations

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run create_migration.go <migration_description>")
		os.Exit(1)
	}

	migrationDescription := strings.Join(os.Args[1:], "_")

	migrationName, structName := createMigrationFile(migrationDescription)
	addMigrationToRegistryFile(migrationName, structName)

	formatMigrationFiles()

	fmt.Println("Created migration:", migrationName)
}

func toPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}

	return strings.Join(parts, "")
}

func createMigrationFile(migrationDescription string) (migrationName string, structName string) {
	migrationDescription = strings.ToLower(strings.ReplaceAll(migrationDescription, " ", "_"))
	timestamp := time.Now().Format("2006_01_02T150405")

	// migration name is of the form e.g 2025-04-29T100421-add-new-users
	migrationName = fmt.Sprintf("%s_%s", timestamp, migrationDescription)
	filePath := fmt.Sprintf("../migrations/%s.go", migrationName)

	// struct name is of the form e.g AddNewUsers
	structName = toPascalCase(migrationDescription)
	content := fmt.Sprintf(migrationTemplate, structName, structName, structName)

	// Ensure migrations directory exists
	if _, err := os.Stat("../migrations"); os.IsNotExist(err) {
		if err := os.Mkdir("migrations", 0755); err != nil {
			fmt.Println("Failed to create migrations directory:", err)
			os.Exit(1)
		}
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		fmt.Println("Failed to write migration file:", err)
		os.Exit(1)
	}

	return migrationName, structName
}

func addMigrationToRegistryFile(migrationName string, structName string) {
	// Path to your registry.go file
	filename := "../registry.go"

	// The lines you want to insert
	registerMigrationLine := fmt.Sprintf(`addMigrationToRegistry("%s", &migrations.%s{})`, migrationName, structName)

	// Read the original file
	input, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the migration registry:", err)
		return
	}
	defer input.Close()

	var outputLines []string

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		// Insert register migration line before //#registerMigration
		if strings.TrimSpace(line) == "//#registerMigration" {
			outputLines = append(outputLines, registerMigrationLine)
		}

		outputLines = append(outputLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the migration registry:", err)
		return
	}

	// Write back to the file. This truncates the file and opens it for writing
	output, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error opening the migration registry for writing:", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, l := range outputLines {
		fmt.Fprintln(writer, l)
	}
	writer.Flush()

	fmt.Println("Registry updated successfully.")
}

func formatMigrationFiles() {
	cmd := exec.Command("go", "fmt", "../")
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Could not format migration files. Please do this manually")
	}

	fmt.Println("Migration files formatted successfully!")
}
