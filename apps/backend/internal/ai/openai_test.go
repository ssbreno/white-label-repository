package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ssbreno/white-label-repository/backend/internal/config"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestOpenAI_ChatCompletion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "Bearer test-key", r.Header.Get("Authorization"))
		assert.Equal(t, "/chat/completions", r.URL.Path)

		resp := openaiResponse{
			ID: "chatcmpl-123",
			Choices: []struct {
				Message struct {
					Role    string `json:"role"`
					Content string `json:"content"`
				} `json:"message"`
			}{
				{Message: struct {
					Role    string `json:"role"`
					Content string `json:"content"`
				}{Role: "assistant", Content: "Hello!"}},
			},
			Usage: struct {
				PromptTokens     int `json:"prompt_tokens"`
				CompletionTokens int `json:"completion_tokens"`
				TotalTokens      int `json:"total_tokens"`
			}{PromptTokens: 10, CompletionTokens: 5, TotalTokens: 15},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	provider := NewOpenAIProvider(config.ProviderConfig{
		APIKey:  "test-key",
		BaseURL: server.URL,
		Enabled: true,
	})

	resp, err := provider.ChatCompletion(context.Background(), models.ChatCompletionRequest{
		Model:    "gpt-4",
		Messages: []models.ChatMessage{{Role: "user", Content: "Hi"}},
	})

	assert.NoError(t, err)
	assert.Equal(t, "chatcmpl-123", resp.ID)
	assert.Equal(t, "Hello!", resp.Message.Content)
	assert.Equal(t, "openai", resp.Provider)
	assert.Equal(t, 15, resp.Usage.TotalTokens)
}

func TestOpenAI_ChatCompletion_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "invalid api key"}`))
	}))
	defer server.Close()

	provider := NewOpenAIProvider(config.ProviderConfig{
		APIKey:  "bad-key",
		BaseURL: server.URL,
		Enabled: true,
	})

	_, err := provider.ChatCompletion(context.Background(), models.ChatCompletionRequest{
		Model:    "gpt-4",
		Messages: []models.ChatMessage{{Role: "user", Content: "Hi"}},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "401")
}

func TestOpenAI_Models(t *testing.T) {
	provider := NewOpenAIProvider(config.ProviderConfig{})
	m := provider.Models()
	assert.Greater(t, len(m), 0)
	for _, model := range m {
		assert.Equal(t, "openai", model.Provider)
		assert.NotEmpty(t, model.ID)
		assert.NotEmpty(t, model.Name)
	}
}

func TestOpenAI_Name(t *testing.T) {
	provider := NewOpenAIProvider(config.ProviderConfig{})
	assert.Equal(t, "openai", provider.Name())
}
