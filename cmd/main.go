package main

import (
	"github.com/babacar-thiam/go-rbac-api/internal/app"
)

func main() {
	// Create and initialize the application
	application := &app.App{}
	application.Initialize()

	// Run the application
	application.Run()
}
