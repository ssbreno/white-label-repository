package services

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
	"github.com/ssbreno/white-label-repository/backend/internal/repositories"
)

// UserService handles business logic for user operations
type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAll returns all users
func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.FindAll(ctx)
}

// GetByID returns a single user
func (s *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

// Create creates a new user
func (s *UserService) Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	return s.repo.Create(ctx, req)
}

// Update updates an existing user
func (s *UserService) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error) {
	return s.repo.Update(ctx, id, req)
}

// Delete deletes a user
func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
