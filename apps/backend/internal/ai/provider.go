package ai

import (
	"context"

	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// Provider defines the contract every AI backend must fulfill
type Provider interface {
	// Name returns the provider identifier (e.g., "openai", "anthropic", "google")
	Name() string

	// ChatCompletion sends a synchronous chat request and returns the full response
	ChatCompletion(ctx context.Context, req models.ChatCompletionRequest) (*models.ChatCompletionResponse, error)

	// StreamCompletion sends a streaming chat request and writes chunks to the channel.
	// The provider MUST close the channel when done.
	StreamCompletion(ctx context.Context, req models.ChatCompletionRequest, chunks chan<- models.StreamChunk) error

	// Models returns the list of models this provider supports
	Models() []models.ModelInfo
}
