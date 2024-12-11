package main

import (
	"log"

	"github.com/subosito/gotenv"

	"github.com/babacar-thiam/go-rbac-api/internal/app"
)

func main() {
	// Load .env file
	err := gotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Create and initialize the application
	application := &app.App{}
	application.Init()

	// Run the application
	application.Run()
}
