package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Role holds the structure of role model
type Role struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RequiredFields checks for missing required fields
func (r *Role) RequiredFields() error {
	fields := map[string]string{
		"name":        r.Name,
		"description": r.Description,
	}

	// iterate through the required fields
	for field, value := range fields {
		if value == "" {
			return fmt.Errorf("%s is required", field)
		}
	}

	return nil
}

// Validate ensures the required fields are non-emtpy
func (r *Role) Validate() error {
	return r.RequiredFields()
}
