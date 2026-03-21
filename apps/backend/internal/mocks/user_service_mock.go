package mocks

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// MockUserService is a mock implementation of UserServiceInterface
type MockUserService struct {
	GetAllFn  func(ctx context.Context) ([]models.User, error)
	GetByIDFn func(ctx context.Context, id string) (*models.User, error)
	CreateFn  func(ctx context.Context, req models.CreateUserRequest) (*models.User, error)
	UpdateFn  func(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error)
	DeleteFn  func(ctx context.Context, id string) error
}

func (m *MockUserService) GetAll(ctx context.Context) ([]models.User, error) {
	return m.GetAllFn(ctx)
}

func (m *MockUserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return m.GetByIDFn(ctx, id)
}

func (m *MockUserService) Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	return m.CreateFn(ctx, req)
}

func (m *MockUserService) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error) {
	return m.UpdateFn(ctx, id, req)
}

func (m *MockUserService) Delete(ctx context.Context, id string) error {
	return m.DeleteFn(ctx, id)
}
