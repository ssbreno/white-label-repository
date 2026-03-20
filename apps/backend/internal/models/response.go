package models

import "time"

// APIResponse is a generic API response wrapper
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// PaginatedResponse wraps a list with pagination metadata
type PaginatedResponse[T any] struct {
	Items      []T   `json:"items"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	TotalPages int   `json:"total_pages"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Version   string            `json:"version"`
	Services  map[string]string `json:"services"`
}

// NewSuccess creates a successful API response
func NewSuccess[T any](data T, message string) APIResponse[T] {
	return APIResponse[T]{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// NewError creates an error API response
func NewError[T any](errMsg string) APIResponse[T] {
	return APIResponse[T]{
		Success: false,
		Error:   errMsg,
	}
}
