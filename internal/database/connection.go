package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

// Connect oppens a database connection
func Connect() (*sql.DB, error) {
	// Load .env file
	err := gotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// Read environment variables
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// DSN(database source name)
	// Database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Test the database connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Connected successfully to database")
	return db, nil
}

// Close closes the database connection
func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("error closing database connection: %v", err)
	}
}
