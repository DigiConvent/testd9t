package install

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/DigiConvent/testd9t/core/env"
	db "github.com/DigiConvent/testd9t/core/pg"
	"github.com/DigiConvent/testd9t/install/version"
)

func MigrateTo(toVersion version.SemVer) {
	conn := db.GetPGDBConnection()
	lastMigration := conn.QueryRow(context.Background(), "SELECT major, minor, patch FROM versions WHERE migrated=true ORDER BY major DESC, minor DESC, patch DESC LIMIT 1")
	var major, minor, patch int
	currentVersion := &version.SemVer{Major: -1, Minor: -1, Patch: -1}

	err := lastMigration.Scan(&currentVersion.Major, &currentVersion.Minor, &currentVersion.Patch)
	if err != nil {
		if strings.Contains(err.Error(), "42P01") {
			fmt.Println("Empty database, starting from scratch.")
		} else {
			fmt.Println("Error fetching last migration:", err)
		}
	} else {
		fmt.Printf("Migrate %s -> %s\n", currentVersion.String(), toVersion.String())
	}

	fmt.Println("From:", major, minor, patch)
	fmt.Println("To  :", toVersion.Major, toVersion.Minor, toVersion.Patch)

	var migrator version.Migrator

	if version.Dev() {
		migrator = version.NewLocalMigrationsMigrator()
	} else {
		migrator = version.NewRemoteMigrationsMigrator()
	}

	versions := migrator.ListMigrationVersions()

	for _, version := range versions {
		if !version.SmallerThan(&toVersion) {
			continue
		}

		for file, contents := range migrator.ListMigrationFiles(&version) {
			fmt.Println("Exec:", version.String(), file)
			_, err := conn.Exec(context.Background(), contents)

			if err != nil {
				fmt.Println("Error executing migration:", err)
				return
			}
		}
		conn.Exec(context.Background(), "INSERT INTO versions (major, minor, patch, migrated) VALUES ($1, $2, $3, true)", version.Major, version.Minor, version.Patch)
		fmt.Println("Done:", version.String())
	}
}

func ResetDB() {
	queries := []string{
		"SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '${DB_NAME}';",
		"DROP DATABASE ${DB_NAME};",
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '${DB_NAME}');",
		"DROP OWNED BY d9t;",
		"REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM d9t;",
		"REVOKE ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public FROM d9t;",
		"REVOKE ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public FROM d9t;",
		"DROP ROLE d9t;",
		"CREATE ROLE d9t WITH LOGIN PASSWORD '${DB_PASSWORD}';",
		"CREATE DATABASE ${DB_NAME} OWNER d9t;",
	}

	for _, query := range queries {
		PsqlCommand(query)
	}
}

func setPGPW() {
	err := os.Setenv("PGPASSWORD", os.Getenv("DB_PASSWORD"))
	if err != nil {
		fmt.Println("Could not set PGPASSWORD environment variable.")
	}
}

func PsqlCommand(command string) {
	setPGPW()
	cmd := exec.Command("sudo", "-u", "postgres", "psql", "-d", "postgres", "-c", env.ReplaceEnvVar(command))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Could not execute query:", command)
		fmt.Println(string(output))
	}
	err = nil
}

func BackupDatabase(filePath string) error {
	dbName := os.Getenv("DB_NAME")
	setPGPW()
	os.Remove(filePath)
	cmd := exec.Command("pg_dump", "-U", "postgres", "-d", dbName, "--data-only", "-f", filePath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing pg_dump: %v\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		return err
	}
	return nil
}

func RestoreDatabase(filePath string) error {
	dbName := os.Getenv("DB_NAME")
	setPGPW()

	cmd := exec.Command("psql", "-U", "postgres", "-d", dbName, "--file="+filePath)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing pg_dump: %v\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		return err
	}

	fmt.Println("Restored data from:", filePath)
	return nil
}
