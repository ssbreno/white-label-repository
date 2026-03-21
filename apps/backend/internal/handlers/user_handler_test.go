package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ssbreno/white-label-repository/backend/internal/mocks"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func setupRouter(svc *mocks.MockUserService) *gin.Engine {
	r := gin.New()
	h := newUserHandler(svc)
	api := r.Group("/api/users")
	{
		api.GET("", h.list)
		api.GET("/:id", h.getByID)
		api.POST("", h.create)
		api.PUT("/:id", h.update)
		api.DELETE("/:id", h.delete)
	}
	return r
}

func testUser() models.User {
	return models.User{
		ID:        "uuid-123",
		Email:     "test@example.com",
		Name:      "Test User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestListUsers(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func(*mocks.MockUserService)
		wantStatus int
		wantLen    int
	}{
		{
			name: "returns users list",
			mockSetup: func(m *mocks.MockUserService) {
				m.GetAllFn = func(ctx context.Context) ([]models.User, error) {
					return []models.User{testUser()}, nil
				}
			},
			wantStatus: http.StatusOK,
			wantLen:    1,
		},
		{
			name: "returns 500 on error",
			mockSetup: func(m *mocks.MockUserService) {
				m.GetAllFn = func(ctx context.Context) ([]models.User, error) {
					return nil, fmt.Errorf("db error")
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mocks.MockUserService{}
			tt.mockSetup(mockSvc)
			router := setupRouter(mockSvc)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/users", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		mockSetup  func(*mocks.MockUserService)
		wantStatus int
	}{
		{
			name: "returns user",
			id:   "uuid-123",
			mockSetup: func(m *mocks.MockUserService) {
				m.GetByIDFn = func(ctx context.Context, id string) (*models.User, error) {
					u := testUser()
					return &u, nil
				}
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "returns 404 when not found",
			id:   "nonexistent",
			mockSetup: func(m *mocks.MockUserService) {
				m.GetByIDFn = func(ctx context.Context, id string) (*models.User, error) {
					return nil, fmt.Errorf("user not found")
				}
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mocks.MockUserService{}
			tt.mockSetup(mockSvc)
			router := setupRouter(mockSvc)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/users/"+tt.id, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name       string
		body       interface{}
		mockSetup  func(*mocks.MockUserService)
		wantStatus int
	}{
		{
			name: "creates user successfully",
			body: models.CreateUserRequest{Email: "new@example.com", Name: "New User"},
			mockSetup: func(m *mocks.MockUserService) {
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
			wantStatus: http.StatusCreated,
		},
		{
			name:       "returns 400 on invalid body",
			body:       map[string]string{"email": "invalid"},
			mockSetup:  func(m *mocks.MockUserService) {},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "returns 500 on service error",
			body: models.CreateUserRequest{Email: "fail@example.com", Name: "Fail User"},
			mockSetup: func(m *mocks.MockUserService) {
				m.CreateFn = func(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
					return nil, fmt.Errorf("duplicate key")
				}
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mocks.MockUserService{}
			tt.mockSetup(mockSvc)
			router := setupRouter(mockSvc)

			body, _ := json.Marshal(tt.body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/users", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	mockSvc := &mocks.MockUserService{
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
	router := setupRouter(mockSvc)

	body, _ := json.Marshal(models.UpdateUserRequest{Email: "updated@example.com", Name: "Updated"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/users/uuid-123", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name       string
		mockSetup  func(*mocks.MockUserService)
		wantStatus int
	}{
		{
			name: "deletes user",
			mockSetup: func(m *mocks.MockUserService) {
				m.DeleteFn = func(ctx context.Context, id string) error {
					return nil
				}
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "returns 404 when not found",
			mockSetup: func(m *mocks.MockUserService) {
				m.DeleteFn = func(ctx context.Context, id string) error {
					return fmt.Errorf("user not found")
				}
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := &mocks.MockUserService{}
			tt.mockSetup(mockSvc)
			router := setupRouter(mockSvc)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/api/users/uuid-123", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}
