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
package %s

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

	migrationName, packageName, structName := createMigrationFile(migrationDescription)

	addMigrationToRegistryFile(migrationName, packageName, structName)
	importMigrationInRegistryFile(packageName, migrationName)
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

func createMigrationFile(migrationDescription string) (migrationName string, packageName string, structName string) {
	migrationDescription = strings.ToLower(strings.ReplaceAll(migrationDescription, " ", "_"))
	timestamp := time.Now().Format("2006_01_02T150405")

	// migration name is of the form e.g 2025_04_29T100421_add_new_users
	migrationName = fmt.Sprintf("%s_%s", timestamp, migrationDescription)
	directoryPath := fmt.Sprintf("../migrations/%s", migrationName)
	filePath := fmt.Sprintf("%s/run.go", directoryPath)

	// package name is of the form e.g add_new_users_2025_04_29T100421
	// We need to have the description first because package names can only begin with letters
	packageName = fmt.Sprintf("%s_%s", migrationDescription, timestamp)
	// struct name is of the form e.g AddNewUsers
	structName = toPascalCase(packageName)

	content := fmt.Sprintf(migrationTemplate, packageName, structName, structName, structName)

	// Create migration directory. This is where run.go will live
	if err := os.Mkdir(directoryPath, 0755); err != nil {
		fmt.Println("Failed to create migration directory:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		fmt.Println("Failed to write migration file:", err)
		os.Exit(1)
	}

	return migrationName, packageName, structName
}

func addMigrationToRegistryFile(migrationName string, packageName string, structName string) {
	// Path to your registry.go file
	filename := "../registry.go"

	// The lines you want to insert
	registerMigrationLine := fmt.Sprintf(`addMigrationToRegistry("%s", &%s.%s{})`, migrationName, packageName, structName)

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

	fmt.Println("Registry updated with migration registration successfully.")
}

func importMigrationInRegistryFile(packageName string, migrationName string) {
	// Path to your registry.go file
	filename := "../registry.go"

	// The lines you want to insert
	importMigrationLine := fmt.Sprintf(`%s "github.com/gbenga504/travel-assistant/utils/migrate/migrations/%s"`, packageName, migrationName)

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

		// Insert import migration line before //#importMigration
		if strings.TrimSpace(line) == "//#importMigration" {
			outputLines = append(outputLines, importMigrationLine)
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

	fmt.Println("Registry updated with migration import successfully")
}

func formatMigrationFiles() {
	cmd := exec.Command("go", "fmt", "../")
	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Could not format migration files. Please do this manually")
	}

	fmt.Println("Migration files formatted successfully!")
}
