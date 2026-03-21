package models

// ChatMessage represents a single message in a conversation
type ChatMessage struct {
	Role    string `json:"role" binding:"required,oneof=system user assistant"`
	Content string `json:"content" binding:"required"`
}

// ChatCompletionRequest is the unified request for all providers
type ChatCompletionRequest struct {
	Provider    string        `json:"provider" binding:"required"`
	Model       string        `json:"model" binding:"required"`
	Messages    []ChatMessage `json:"messages" binding:"required,min=1"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature *float64      `json:"temperature,omitempty"`
	Stream      bool          `json:"stream,omitempty"`
	TopP        *float64      `json:"top_p,omitempty"`
	Stop        []string      `json:"stop,omitempty"`
}

// ChatCompletionResponse is the unified response from all providers
type ChatCompletionResponse struct {
	ID       string      `json:"id"`
	Provider string      `json:"provider"`
	Model    string      `json:"model"`
	Message  ChatMessage `json:"message"`
	Usage    TokenUsage  `json:"usage"`
}

// TokenUsage tracks token consumption
type TokenUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// StreamChunk represents a single SSE chunk during streaming
type StreamChunk struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Delta    string `json:"delta"`
	Done     bool   `json:"done"`
}

// ModelInfo describes an available model
type ModelInfo struct {
	ID            string  `json:"id"`
	Provider      string  `json:"provider"`
	Name          string  `json:"name"`
	MaxTokens     int     `json:"max_tokens"`
	InputPricing  float64 `json:"input_pricing_per_1k"`
	OutputPricing float64 `json:"output_pricing_per_1k"`
}
