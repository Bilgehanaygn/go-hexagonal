package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // ← or mysql, sqlite3, etc, matching your DB_URL
	_ "github.com/golang-migrate/migrate/v4/source/file"       // ← register the “file://” source driver
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("Port not found in the env fil")
	}
	fmt.Println("Listening on port:", port)

	m, err := migrate.New(
		"file://db/migrations",
		dbUrl,
	)

	if err != nil {
		log.Fatalf("failed to initialize migrations: %v", err)
	}

	// this is what actually *runs* the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}
}
