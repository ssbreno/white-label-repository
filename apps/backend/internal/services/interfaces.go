package services

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// UserServiceInterface defines the contract for user business logic
type UserServiceInterface interface {
	GetAll(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error)
	Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
