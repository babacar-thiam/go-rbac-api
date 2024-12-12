package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// User holds the structure of user model
type User struct {
	ID           uuid.UUID `json:"id"`
	Role         string    `json:"role"`
	FullName     string    `json:"full_name"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// RequiredFields check for missing required fields
func (u *User) RequiredFields() error {
	fields := map[string]string{
		"full name":     u.FullName,
		"email address": u.EmailAddress,
		"phone number":  u.PhoneNumber,
		"password":      u.PasswordHash,
	}

	// iterate through the fields
	for field, value := range fields {
		if value == "" {
			return fmt.Errorf("%s is required", field)
		}
	}

	return nil
}

// Validate ensures the required fields are non-emtpy
func (u *User) Validate() error {
	return u.RequiredFields()
}
