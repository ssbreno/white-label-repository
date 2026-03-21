package mocks

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// MockUserRepository is a mock implementation of UserRepositoryInterface
type MockUserRepository struct {
	FindAllFn  func(ctx context.Context) ([]models.User, error)
	FindByIDFn func(ctx context.Context, id string) (*models.User, error)
	CreateFn   func(ctx context.Context, req models.CreateUserRequest) (*models.User, error)
	UpdateFn   func(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error)
	DeleteFn   func(ctx context.Context, id string) error
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	return m.FindAllFn(ctx)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	return m.FindByIDFn(ctx, id)
}

func (m *MockUserRepository) Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	return m.CreateFn(ctx, req)
}

func (m *MockUserRepository) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error) {
	return m.UpdateFn(ctx, id, req)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}
