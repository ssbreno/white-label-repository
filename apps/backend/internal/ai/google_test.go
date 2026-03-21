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

func TestGoogle_ChatCompletion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Contains(t, r.URL.Path, "generateContent")
		assert.Equal(t, "test-key", r.URL.Query().Get("key"))

		resp := geminiResponse{
			Candidates: []struct {
				Content struct {
					Parts []struct {
						Text string `json:"text"`
					} `json:"parts"`
					Role string `json:"role"`
				} `json:"content"`
			}{
				{Content: struct {
					Parts []struct {
						Text string `json:"text"`
					} `json:"parts"`
					Role string `json:"role"`
				}{
					Parts: []struct {
						Text string `json:"text"`
					}{{Text: "Hello!"}},
					Role: "model",
				}},
			},
			UsageMetadata: struct {
				PromptTokenCount     int `json:"promptTokenCount"`
				CandidatesTokenCount int `json:"candidatesTokenCount"`
				TotalTokenCount      int `json:"totalTokenCount"`
			}{PromptTokenCount: 10, CandidatesTokenCount: 5, TotalTokenCount: 15},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	provider := NewGoogleProvider(config.ProviderConfig{
		APIKey:  "test-key",
		BaseURL: server.URL,
		Enabled: true,
	})

	resp, err := provider.ChatCompletion(context.Background(), models.ChatCompletionRequest{
		Model:    "gemini-1.5-pro",
		Messages: []models.ChatMessage{{Role: "user", Content: "Hi"}},
	})

	assert.NoError(t, err)
	assert.Equal(t, "Hello!", resp.Message.Content)
	assert.Equal(t, "google", resp.Provider)
	assert.Equal(t, 15, resp.Usage.TotalTokens)
}

func TestGoogle_Models(t *testing.T) {
	provider := NewGoogleProvider(config.ProviderConfig{})
	m := provider.Models()
	assert.Greater(t, len(m), 0)
	for _, model := range m {
		assert.Equal(t, "google", model.Provider)
	}
}

func TestGoogle_Name(t *testing.T) {
	provider := NewGoogleProvider(config.ProviderConfig{})
	assert.Equal(t, "google", provider.Name())
}
