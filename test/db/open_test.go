package test

import (
	"log"
	"os"
	"testing"

	"github.com/subosito/gotenv"

	"github.com/babacar-thiam/go-rbac-api/internal/db"
)

func TestMain(m *testing.M) {
	// Load the .env.test file
	err := gotenv.Load("../../.env.test")
	if err != nil {
		log.Fatalf("error loading .env file")
	}

	// Run tests
	os.Exit(m.Run())
}

func TestOpen(t *testing.T) {
	// Open the database connection
	db, err := db.OpenDB()
	if err != nil {
		t.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	// Perform a simple ping to verify the connection
	if err := db.Ping(); err != nil {
		t.Errorf("database connection not healthy: %v", err)
	}
}
