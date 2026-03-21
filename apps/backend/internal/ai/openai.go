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

// OpenAIProvider implements the Provider interface for OpenAI
type OpenAIProvider struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider(cfg config.ProviderConfig) *OpenAIProvider {
	return &OpenAIProvider{
		apiKey:  cfg.APIKey,
		baseURL: cfg.BaseURL,
		client:  &http.Client{Timeout: 120 * time.Second},
	}
}

func (p *OpenAIProvider) Name() string { return "openai" }

// openaiRequest is the OpenAI API request format
type openaiRequest struct {
	Model       string              `json:"model"`
	Messages    []openaiMessage     `json:"messages"`
	MaxTokens   int                 `json:"max_tokens,omitempty"`
	Temperature *float64            `json:"temperature,omitempty"`
	TopP        *float64            `json:"top_p,omitempty"`
	Stop        []string            `json:"stop,omitempty"`
	Stream      bool                `json:"stream,omitempty"`
}

type openaiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// openaiResponse is the OpenAI API response format
type openaiResponse struct {
	ID      string `json:"id"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type openaiStreamChunk struct {
	ID      string `json:"id"`
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason *string `json:"finish_reason"`
	} `json:"choices"`
}

func (p *OpenAIProvider) ChatCompletion(ctx context.Context, req models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	oaiReq := p.toOpenAIRequest(req)
	oaiReq.Stream = false

	body, err := json.Marshal(oaiReq)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("openai API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var oaiResp openaiResponse
	if err := json.NewDecoder(resp.Body).Decode(&oaiResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(oaiResp.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}

	return &models.ChatCompletionResponse{
		ID:       oaiResp.ID,
		Provider: "openai",
		Model:    req.Model,
		Message: models.ChatMessage{
			Role:    oaiResp.Choices[0].Message.Role,
			Content: oaiResp.Choices[0].Message.Content,
		},
		Usage: models.TokenUsage{
			PromptTokens:     oaiResp.Usage.PromptTokens,
			CompletionTokens: oaiResp.Usage.CompletionTokens,
			TotalTokens:      oaiResp.Usage.TotalTokens,
		},
	}, nil
}

func (p *OpenAIProvider) StreamCompletion(ctx context.Context, req models.ChatCompletionRequest, chunks chan<- models.StreamChunk) error {
	defer close(chunks)

	oaiReq := p.toOpenAIRequest(req)
	oaiReq.Stream = true

	body, err := json.Marshal(oaiReq)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/chat/completions", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("openai API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			chunks <- models.StreamChunk{Provider: "openai", Model: req.Model, Done: true}
			return nil
		}

		var chunk openaiStreamChunk
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			continue
		}
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			chunks <- models.StreamChunk{
				ID:       chunk.ID,
				Provider: "openai",
				Model:    req.Model,
				Delta:    chunk.Choices[0].Delta.Content,
			}
		}
	}

	return scanner.Err()
}

func (p *OpenAIProvider) Models() []models.ModelInfo {
	return []models.ModelInfo{
		{ID: "gpt-4o", Provider: "openai", Name: "GPT-4o", MaxTokens: 128000, InputPricing: 0.0025, OutputPricing: 0.01},
		{ID: "gpt-4o-mini", Provider: "openai", Name: "GPT-4o Mini", MaxTokens: 128000, InputPricing: 0.00015, OutputPricing: 0.0006},
		{ID: "gpt-4-turbo", Provider: "openai", Name: "GPT-4 Turbo", MaxTokens: 128000, InputPricing: 0.01, OutputPricing: 0.03},
		{ID: "gpt-3.5-turbo", Provider: "openai", Name: "GPT-3.5 Turbo", MaxTokens: 16385, InputPricing: 0.0005, OutputPricing: 0.0015},
	}
}

func (p *OpenAIProvider) toOpenAIRequest(req models.ChatCompletionRequest) openaiRequest {
	msgs := make([]openaiMessage, len(req.Messages))
	for i, m := range req.Messages {
		msgs[i] = openaiMessage{Role: m.Role, Content: m.Content}
	}
	return openaiRequest{
		Model:       req.Model,
		Messages:    msgs,
		MaxTokens:   req.MaxTokens,
		Temperature: req.Temperature,
		TopP:        req.TopP,
		Stop:        req.Stop,
	}
}
