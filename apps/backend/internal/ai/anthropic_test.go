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

func TestAnthropic_ChatCompletion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "test-key", r.Header.Get("x-api-key"))
		assert.Equal(t, "2023-06-01", r.Header.Get("anthropic-version"))
		assert.Equal(t, "/v1/messages", r.URL.Path)

		// Verify system message is extracted
		var req anthropicRequest
		json.NewDecoder(r.Body).Decode(&req)
		assert.Equal(t, "You are helpful", req.System)
		assert.Len(t, req.Messages, 1)
		assert.Equal(t, "user", req.Messages[0].Role)

		resp := anthropicResponse{
			ID: "msg-123",
			Content: []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}{
				{Type: "text", Text: "Hello!"},
			},
			Usage: struct {
				InputTokens  int `json:"input_tokens"`
				OutputTokens int `json:"output_tokens"`
			}{InputTokens: 10, OutputTokens: 5},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	provider := NewAnthropicProvider(config.ProviderConfig{
		APIKey:  "test-key",
		BaseURL: server.URL,
		Enabled: true,
	})

	resp, err := provider.ChatCompletion(context.Background(), models.ChatCompletionRequest{
		Model: "claude-3-5-sonnet-20241022",
		Messages: []models.ChatMessage{
			{Role: "system", Content: "You are helpful"},
			{Role: "user", Content: "Hi"},
		},
		MaxTokens: 1024,
	})

	assert.NoError(t, err)
	assert.Equal(t, "msg-123", resp.ID)
	assert.Equal(t, "Hello!", resp.Message.Content)
	assert.Equal(t, "anthropic", resp.Provider)
	assert.Equal(t, 15, resp.Usage.TotalTokens)
}

func TestAnthropic_Models(t *testing.T) {
	provider := NewAnthropicProvider(config.ProviderConfig{})
	m := provider.Models()
	assert.Greater(t, len(m), 0)
	for _, model := range m {
		assert.Equal(t, "anthropic", model.Provider)
	}
}

func TestAnthropic_Name(t *testing.T) {
	provider := NewAnthropicProvider(config.ProviderConfig{})
	assert.Equal(t, "anthropic", provider.Name())
}
