package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("Port not found in the env fil")
	}
	fmt.Print("Port:", port)

	  m, err := migrate.New(
    "file://path/to/migrations",
    "postgres://user:pass@localhost:5432/dbname?sslmode=disable",
  )

    if err != nil {
    log.Fatalf("failed to initialize migrations: %v", err)
  }

  // this is what actually *runs* the migrations
  if err := m.Up(); err != nil && err != migrate.ErrNoChange {
    log.Fatalf("migration failed: %v", err)
  }
}