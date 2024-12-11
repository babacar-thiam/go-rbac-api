package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	// Initialize role repository
	roleRepo := repositories.NewRepository(db)

	// Begin database transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		// Ensure the rollback in case of failure
		if rollbackErr := tx.Rollback(); rollbackErr != nil &&
			!errors.Is(rollbackErr, sql.ErrTxDone) {
			log.Printf("failed to rollback transaction: %v", rollbackErr)
		}
	}()

	// Use context with timeout for queries
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("starting role seeding with %d default roles", len(defaultRoles))

	// Iterate through de default roles
	for _, role := range defaultRoles {
		// Check if role already exists
		existingRole, getErr := roleRepo.FindRoleByName(ctx, role.Name)

		// If we got ErrNoRows, it means the role doesn't exist and we should create it
		if errors.Is(getErr, sql.ErrNoRows) {
			// Role does not exist, create it
			if createErr := roleRepo.SaveRole(ctx, role); createErr != nil {
				return fmt.Errorf("failed to create role %s: %v", role.Name, createErr)
			}
			log.Printf("created role: %s", role.Name)
			continue
		}

		// If we get any other error, return it
		if getErr != nil {
			return fmt.Errorf("failed to check role %s: %v", role.Name, getErr)
		}

		// If we get here, the role exists
		log.Printf("role %s already exists, skipping...", existingRole.Name)
	}

	// Commit the transaction
	if commitErr := tx.Commit(); commitErr != nil {
		return fmt.Errorf("faild to commit transaction: %v", commitErr)
	}

	log.Printf("role seeding completed successfully")
	return nil
}
