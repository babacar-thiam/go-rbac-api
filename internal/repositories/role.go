package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

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
func (r *Repository) SaveRole(role models.Role) error {
	// Define SQL query to insert role to the database
	query := "insert into roles (id, name, description, created_at, updated_at) values (?, ?, ?, ?, ?)"

	// Execute the query with context
	_, err := r.db.Exec(
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
func (r *Repository) FindAllRoles() ([]models.Role, error) {
	// Define SQL query to retrieve all roles
	query := "select id, name, description, created_at, updated_at from roles"

	// Execute the query to fetch all roles
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error quering roles: %v", err)
	}
	defer rows.Close()

	// Create a variable that holds all roles data
	var roles []models.Role

	for rows.Next() {
		var role models.Role
		// Scan each role into a Role struct
		scanErr := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		)
		if scanErr != nil {
			return nil, fmt.Errorf("error scanning rows: %v", scanErr)
		}
		roles = append(roles, role)
	}

	// Check for errors encounter during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return roles, nil
}

// FindRoleByID retrieves a single role by its ID
func (r Repository) FindRoleByID(id uuid.UUID) (*models.Role, error) {
	// Define SQL query to retrieve role by its name
	query := "select id, name, description, created_at, updated_at from roles where id = ?"

	// Create a variable to hold the role data
	var role models.Role

	// Execute the query and map the result to the Role struct
	err := r.db.QueryRow(query, id).
		Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No role found
		}
		return nil, fmt.Errorf("error scanning role row: %v", err) // Other errors
	}

	// Return the role as a pointer
	return &role, nil
}

// FindRoleByName retrieves a single role by its name
func (r Repository) FindRoleByName(name string) (*models.Role, error) {
	// Define SQL query to retrieve role by its name
	query := "select id, name, description, created_at, updated_at from roles where name = ?"

	// Create a variable to hold the role data
	var role models.Role

	// Execute the query and map the result to the Role struct
	err := r.db.QueryRow(query, name).
		Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No role found
		}
		return nil, fmt.Errorf("error scanning role row: %v", err) // Other errors
	}

	// Return the role as a pointer
	return &role, nil
}
