package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/babacar-thiam/go-rbac-api/internal/models"
	"github.com/babacar-thiam/go-rbac-api/internal/repositories"
)

// SeedRoles creates initial roles if they don't exist
func SeedRole(db *sql.DB) error {
	// Define default roles
	defaultRoles := []models.Role{
		{
			ID: uuid.New(), Name: "ADMIN",
			Description: "Administrator role",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "CLIENT",
			Description: "CLient role",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New(),
			Name:        "PROVIDER",
			Description: "Provider role",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Intialize roles repository
	roleRepo := repositories.NewRepository(db)

	// Start the transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check and create each role
	for _, role := range defaultRoles {
		// Check if the role already exists
		existingRole, err := roleRepo.FindRoleByName(role.Name)
		if err != nil {
			log.Printf("error checking role %s: %v", role.Name, err)
			return err
		}

		if existingRole != nil {
			log.Printf("role %s already exists, skipping...", role.Name)
			continue
		}

		// Create role if it does not exist
		err = roleRepo.SaveRole(role)
		if err != nil {
			return err
		}
		log.Printf("created role: %s", role.Name)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	log.Println("role seeding completed successfully")
	return nil
}
