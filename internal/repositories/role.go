package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/babacar-thiam/go-rbac-api/internal/models"
)

// Repository handles database operations for roles
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new role repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// SaveRole inserts a new role to the database
func (r *Repository) SaveRole(ctx context.Context, role models.Role) error {
	// Define SQL query to insert role to the database
	query := "insert into roles (id, name, description, created_at, updated_at) values (?, ?, ?, ?, ?)"

	// Execute the query with context
	_, err := r.db.ExecContext(
		ctx,
		query,
		role.ID,
		role.Name,
		role.Description,
		role.CreatedAt,
		role.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("error saving role: %v", err)
	}
	return nil
}

// FindAllRoles retrieve all the roles from the database

// FindRoleByID retrieves a single role by its ID

// FindRoleByName retrieves a single role by its name
func (r Repository) FindRoleByName(ctx context.Context, name string) (models.Role, error) {
	// Define SQL query to retrieve role by its name
	query := "select id, name, description, created_at, updated_at from roles where name = ?"

	// Execute the query with context and fetch a single role
	row := r.db.QueryRowContext(ctx, query, name)

	// Create a variable that holds the role data
	var role models.Role

	// Map the retrieved row to the Role struct using Sccan
	err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Role{}, sql.ErrNoRows
		}
		return models.Role{}, fmt.Errorf("error scanning role row: %v", err)
	}

	return role, nil
}
