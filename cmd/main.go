package main

import (
	"database/sql"
	"embed"
	"fmt"
	"psql_migrations/internal/migrator"
)

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func main() {
	// --- (1) ----
	// Recover Migrator
	migrator := migrator.MustGetNewMigrator(MigrationsFS, migrationsDir)

	// --- (2) ----
	// Get the DB instance
	connectionStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	// --- (2) ----
	// Apply migrations
	err = migrator.ApplyMigrations(conn)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Migrations applied!!")
}
