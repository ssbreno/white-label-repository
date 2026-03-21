package repositories

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// UserRepositoryInterface defines the contract for user data access
type UserRepositoryInterface interface {
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error)
	Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
