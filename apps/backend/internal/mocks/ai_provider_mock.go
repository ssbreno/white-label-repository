package mocks

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// MockAIProvider is a mock implementation of ai.Provider
type MockAIProvider struct {
	NameFn             func() string
	ChatCompletionFn   func(ctx context.Context, req models.ChatCompletionRequest) (*models.ChatCompletionResponse, error)
	StreamCompletionFn func(ctx context.Context, req models.ChatCompletionRequest, chunks chan<- models.StreamChunk) error
	ModelsFn           func() []models.ModelInfo
}

func (m *MockAIProvider) Name() string {
	return m.NameFn()
}

func (m *MockAIProvider) ChatCompletion(ctx context.Context, req models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	return m.ChatCompletionFn(ctx, req)
}

func (m *MockAIProvider) StreamCompletion(ctx context.Context, req models.ChatCompletionRequest, chunks chan<- models.StreamChunk) error {
	return m.StreamCompletionFn(ctx, req, chunks)
}

func (m *MockAIProvider) Models() []models.ModelInfo {
	return m.ModelsFn()
}
