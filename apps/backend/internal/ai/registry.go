package ai

import (
	"fmt"
	"sync"

	"github.com/ssbreno/white-label-repository/backend/internal/config"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// Registry manages available AI providers
type Registry struct {
	mu        sync.RWMutex
	providers map[string]Provider
}

// NewRegistry creates a registry and auto-registers enabled providers
func NewRegistry(cfg config.AIConfig) *Registry {
	r := &Registry{providers: make(map[string]Provider)}

	if cfg.OpenAI.Enabled {
		r.Register(NewOpenAIProvider(cfg.OpenAI))
	}
	if cfg.Anthropic.Enabled {
		r.Register(NewAnthropicProvider(cfg.Anthropic))
	}
	if cfg.Google.Enabled {
		r.Register(NewGoogleProvider(cfg.Google))
	}

	return r
}

// Register adds a provider to the registry
func (r *Registry) Register(p Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.providers[p.Name()] = p
}

// Get returns a provider by name
func (r *Registry) Get(name string) (Provider, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.providers[name]
	if !ok {
		return nil, fmt.Errorf("provider %q not registered", name)
	}
	return p, nil
}

// AllModels returns models from all registered providers
func (r *Registry) AllModels() []models.ModelInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var all []models.ModelInfo
	for _, p := range r.providers {
		all = append(all, p.Models()...)
	}
	return all
}

// ProviderNames returns the names of all registered providers
func (r *Registry) ProviderNames() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.providers))
	for name := range r.providers {
		names = append(names, name)
	}
	return names
}
