package ai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ssbreno/white-label-repository/backend/internal/config"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// AnthropicProvider implements the Provider interface for Anthropic
type AnthropicProvider struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewAnthropicProvider creates a new Anthropic provider
func NewAnthropicProvider(cfg config.ProviderConfig) *AnthropicProvider {
	return &AnthropicProvider{
		apiKey:  cfg.APIKey,
		baseURL: cfg.BaseURL,
		client:  &http.Client{Timeout: 120 * time.Second},
	}
}

func (p *AnthropicProvider) Name() string { return "anthropic" }

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system,omitempty"`
	Messages  []anthropicMessage `json:"messages"`
	Stream    bool               `json:"stream,omitempty"`
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	ID      string `json:"id"`
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	Usage struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
}

type anthropicStreamEvent struct {
	Type  string `json:"type"`
	Delta *struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"delta,omitempty"`
}

func (p *AnthropicProvider) ChatCompletion(ctx context.Context, req models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	antReq := p.toAnthropicRequest(req)
	antReq.Stream = false

	body, err := json.Marshal(antReq)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/v1/messages", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", p.apiKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("anthropic API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var antResp anthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&antResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	content := ""
	for _, c := range antResp.Content {
		if c.Type == "text" {
			content += c.Text
		}
	}

	return &models.ChatCompletionResponse{
		ID:       antResp.ID,
		Provider: "anthropic",
		Model:    req.Model,
		Message: models.ChatMessage{
			Role:    "assistant",
			Content: content,
		},
		Usage: models.TokenUsage{
			PromptTokens:     antResp.Usage.InputTokens,
			CompletionTokens: antResp.Usage.OutputTokens,
			TotalTokens:      antResp.Usage.InputTokens + antResp.Usage.OutputTokens,
		},
	}, nil
}

func (p *AnthropicProvider) StreamCompletion(ctx context.Context, req models.ChatCompletionRequest, chunks chan<- models.StreamChunk) error {
	defer close(chunks)

	antReq := p.toAnthropicRequest(req)
	antReq.Stream = true

	body, err := json.Marshal(antReq)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/v1/messages", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", p.apiKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("anthropic API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")

		var event anthropicStreamEvent
		if err := json.Unmarshal([]byte(data), &event); err != nil {
			continue
		}

		switch event.Type {
		case "content_block_delta":
			if event.Delta != nil && event.Delta.Text != "" {
				chunks <- models.StreamChunk{
					Provider: "anthropic",
					Model:    req.Model,
					Delta:    event.Delta.Text,
				}
			}
		case "message_stop":
			chunks <- models.StreamChunk{Provider: "anthropic", Model: req.Model, Done: true}
			return nil
		}
	}

	return scanner.Err()
}

func (p *AnthropicProvider) Models() []models.ModelInfo {
	return []models.ModelInfo{
		{ID: "claude-sonnet-4-20250514", Provider: "anthropic", Name: "Claude Sonnet 4", MaxTokens: 200000, InputPricing: 0.003, OutputPricing: 0.015},
		{ID: "claude-3-5-sonnet-20241022", Provider: "anthropic", Name: "Claude 3.5 Sonnet", MaxTokens: 200000, InputPricing: 0.003, OutputPricing: 0.015},
		{ID: "claude-3-5-haiku-20241022", Provider: "anthropic", Name: "Claude 3.5 Haiku", MaxTokens: 200000, InputPricing: 0.0008, OutputPricing: 0.004},
	}
}

func (p *AnthropicProvider) toAnthropicRequest(req models.ChatCompletionRequest) anthropicRequest {
	var system string
	var msgs []anthropicMessage

	for _, m := range req.Messages {
		if m.Role == "system" {
			system = m.Content
			continue
		}
		msgs = append(msgs, anthropicMessage{Role: m.Role, Content: m.Content})
	}

	maxTokens := req.MaxTokens
	if maxTokens == 0 {
		maxTokens = 4096
	}

	return anthropicRequest{
		Model:     req.Model,
		MaxTokens: maxTokens,
		System:    system,
		Messages:  msgs,
	}
}
