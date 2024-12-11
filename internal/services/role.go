package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/babacar-thiam/go-rbac-api/internal/models"
	"github.com/babacar-thiam/go-rbac-api/internal/repositories"
)

// Service handles the business logic of roles
type Service struct {
	repo *repositories.Repository
}

// NewService creates a new role Service
func NewService(repo *repositories.Repository) *Service {
	return &Service{repo: repo}
}

// AddRole creates a new role
func (s *Service) AddRole(ctx context.Context, name, description string) error {
	role := models.Role{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return s.repo.SaveRole(ctx, role)
}
