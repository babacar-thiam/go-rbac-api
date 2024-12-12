package app

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/babacar-thiam/go-rbac-api/internal/db"
	"github.com/babacar-thiam/go-rbac-api/internal/handlers"
	"github.com/babacar-thiam/go-rbac-api/internal/repositories"
	"github.com/babacar-thiam/go-rbac-api/internal/services"
)

type App struct {
	db     *sql.DB
	Router *mux.Router
}

// Init sets up the database connection
func (a *App) Init() {
	var err error
	a.db, err = db.OpenDB()
	if err != nil {
		log.Fatalf("failed to intialize database connection: %v", err)
	}

	// Seed roles
	if seedErr := db.SeedRole(a.db); seedErr != nil {
		log.Fatalf("failed to seed roles: %v", seedErr)
	}

	// Initialize router
	a.Router = mux.NewRouter()

	// Set up API routes with prefix "api/v1"
	router := a.Router.PathPrefix("/api/v1").Subrouter()

	// Initialize role repository, service and handler
	roleRepo := repositories.NewRepository(a.db)
	roleService := services.NewService(roleRepo)
	roleHandler := handlers.NewHandler(roleService)

	// Define routes for roles
	router.HandleFunc("/roles", roleHandler.GetRoles).Methods("GET")
}

// Run starts the application
func (a *App) Run() {
	defer db.CloseDB(a.db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("starting server port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
