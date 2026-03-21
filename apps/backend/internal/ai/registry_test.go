package ai

import (
	"testing"

	"github.com/ssbreno/white-label-repository/backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewRegistry_NoProviders(t *testing.T) {
	cfg := config.AIConfig{}
	r := NewRegistry(cfg)

	assert.Empty(t, r.AllModels())
	assert.Empty(t, r.ProviderNames())
}

func TestNewRegistry_WithOpenAI(t *testing.T) {
	cfg := config.AIConfig{
		OpenAI: config.ProviderConfig{
			APIKey:  "test-key",
			BaseURL: "https://api.openai.com/v1",
			Enabled: true,
		},
	}
	r := NewRegistry(cfg)

	names := r.ProviderNames()
	assert.Contains(t, names, "openai")

	provider, err := r.Get("openai")
	assert.NoError(t, err)
	assert.Equal(t, "openai", provider.Name())
}

func TestRegistry_GetUnknownProvider(t *testing.T) {
	r := NewRegistry(config.AIConfig{})

	_, err := r.Get("unknown")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not registered")
}

func TestRegistry_AllModels(t *testing.T) {
	cfg := config.AIConfig{
		OpenAI: config.ProviderConfig{
			APIKey:  "test-key",
			BaseURL: "https://api.openai.com/v1",
			Enabled: true,
		},
		Anthropic: config.ProviderConfig{
			APIKey:  "test-key",
			BaseURL: "https://api.anthropic.com",
			Enabled: true,
		},
	}
	r := NewRegistry(cfg)

	allModels := r.AllModels()
	assert.Greater(t, len(allModels), 0)

	// Should have models from both providers
	hasOpenAI := false
	hasAnthropic := false
	for _, m := range allModels {
		if m.Provider == "openai" {
			hasOpenAI = true
		}
		if m.Provider == "anthropic" {
			hasAnthropic = true
		}
	}
	assert.True(t, hasOpenAI)
	assert.True(t, hasAnthropic)
}
