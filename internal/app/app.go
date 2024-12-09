package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/babacar-thiam/go-rbac-api/internal/database"
)

type App struct {
	db     *sql.DB
	Router *mux.Router
}

// Initialize sets up the database connection
func (a *App) Initialize() {
	var err error
	a.db, err = database.Connect()
	if err != nil {
		log.Fatalf("failed to intialize database connection: %v", err)
	}

	// Initialize router
	a.Router = mux.NewRouter()
}

// Run starts the application
func (a *App) Run() {
	defer database.Close(a.db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("starting server port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
