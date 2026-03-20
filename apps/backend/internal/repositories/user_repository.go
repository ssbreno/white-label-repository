package repositories

import (
	"context"
	"fmt"

	"github.com/ssbreno/white-label-repository/backend/internal/database"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *database.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindAll returns all users
func (r *UserRepository) FindAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.Pool.Query(ctx, `
		SELECT id, email, name, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// FindByID returns a single user by ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var u models.User
	err := r.db.Pool.QueryRow(ctx, `
		SELECT id, email, name, created_at, updated_at
		FROM users WHERE id = $1
	`, id).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &u, nil
}

// Create inserts a new user
func (r *UserRepository) Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	var u models.User
	err := r.db.Pool.QueryRow(ctx, `
		INSERT INTO users (email, name)
		VALUES ($1, $2)
		RETURNING id, email, name, created_at, updated_at
	`, req.Email, req.Name).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &u, nil
}

// Update modifies an existing user
func (r *UserRepository) Update(ctx context.Context, id string, req models.UpdateUserRequest) (*models.User, error) {
	var u models.User
	err := r.db.Pool.QueryRow(ctx, `
		UPDATE users
		SET email = COALESCE(NULLIF($1, ''), email),
		    name  = COALESCE(NULLIF($2, ''), name),
		    updated_at = NOW()
		WHERE id = $3
		RETURNING id, email, name, created_at, updated_at
	`, req.Email, req.Name, id).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &u, nil
}

// Delete removes a user by ID
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.Pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
