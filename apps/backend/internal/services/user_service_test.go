package services

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ssbreno/white-label-repository/backend/internal/mocks"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func newTestUser() models.User {
	return models.User{
		ID:        "test-uuid-123",
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestUserService_GetAll(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocks.MockUserRepository)
		wantLen   int
		wantErr   bool
	}{
		{
			name: "returns users successfully",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.FindAllFn = func(ctx context.Context) ([]models.User, error) {
					return []models.User{newTestUser()}, nil
				}
			},
			wantLen: 1,
		},
		{
			name: "returns empty list",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.FindAllFn = func(ctx context.Context) ([]models.User, error) {
					return []models.User{}, nil
				}
			},
			wantLen: 0,
		},
		{
			name: "returns error on repository failure",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.FindAllFn = func(ctx context.Context) ([]models.User, error) {
					return nil, fmt.Errorf("db error")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockUserRepository{}
			tt.mockSetup(mockRepo)
			svc := NewUserService(mockRepo)

			users, err := svc.GetAll(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, users, tt.wantLen)
		})
	}
}

func TestUserService_GetByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*mocks.MockUserRepository)
		wantErr   bool
	}{
		{
			name: "returns user by ID",
			id:   "test-uuid-123",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.FindByIDFn = func(ctx context.Context, id string) (*models.User, error) {
					u := newTestUser()
					return &u, nil
				}
			},
		},
		{
			name: "returns error when not found",
			id:   "nonexistent",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.FindByIDFn = func(ctx context.Context, id string) (*models.User, error) {
					return nil, fmt.Errorf("user not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockUserRepository{}
			tt.mockSetup(mockRepo)
			svc := NewUserService(mockRepo)

			user, err := svc.GetByID(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.Equal(t, "test-uuid-123", user.ID)
		})
	}
}

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name      string
		req       models.CreateUserRequest
		mockSetup func(*mocks.MockUserRepository)
		wantErr   bool
	}{
		{
			name: "creates user successfully",
			req:  models.CreateUserRequest{Email: "new@example.com", Name: "New User"},
			mockSetup: func(m *mocks.MockUserRepository) {
				m.CreateFn = func(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
					u := models.User{
						ID:        "new-uuid",
						Email:     req.Email,
						Name:      req.Name,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}
					return &u, nil
				}
			},
		},
		{
			name: "returns error on duplicate email",
			req:  models.CreateUserRequest{Email: "dup@example.com", Name: "Dup User"},
			mockSetup: func(m *mocks.MockUserRepository) {
				m.CreateFn = func(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
					return nil, fmt.Errorf("duplicate key")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockUserRepository{}
			tt.mockSetup(mockRepo)
			svc := NewUserService(mockRepo)

			user, err := svc.Create(context.Background(), tt.req)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.req.Email, user.Email)
			assert.Equal(t, tt.req.Name, user.Name)
		})
	}
}

func TestUserService_Update(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		UpdateFn: func(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error) {
			u := models.User{
				ID:        id,
				Email:     req.Email,
				Name:      req.Name,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			return &u, nil
		},
	}
	svc := NewUserService(mockRepo)

	user, err := svc.Update(context.Background(), "test-uuid", models.UpdateUserRequest{
		Email: "updated@example.com",
		Name:  "Updated",
	})
	assert.NoError(t, err)
	assert.Equal(t, "updated@example.com", user.Email)
}

func TestUserService_Delete(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*mocks.MockUserRepository)
		wantErr   bool
	}{
		{
			name: "deletes user successfully",
			id:   "test-uuid",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.DeleteFn = func(ctx context.Context, id string) error {
					return nil
				}
			},
		},
		{
			name: "returns error when not found",
			id:   "nonexistent",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.DeleteFn = func(ctx context.Context, id string) error {
					return fmt.Errorf("user not found")
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.MockUserRepository{}
			tt.mockSetup(mockRepo)
			svc := NewUserService(mockRepo)

			err := svc.Delete(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
