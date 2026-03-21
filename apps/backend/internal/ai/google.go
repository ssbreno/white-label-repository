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

// GoogleProvider implements the Provider interface for Google Gemini
type GoogleProvider struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewGoogleProvider creates a new Google Gemini provider
func NewGoogleProvider(cfg config.ProviderConfig) *GoogleProvider {
	return &GoogleProvider{
		apiKey:  cfg.APIKey,
		baseURL: cfg.BaseURL,
		client:  &http.Client{Timeout: 120 * time.Second},
	}
}

func (p *GoogleProvider) Name() string { return "google" }

type geminiRequest struct {
	Contents         []geminiContent         `json:"contents"`
	GenerationConfig *geminiGenerationConfig `json:"generationConfig,omitempty"`
}

type geminiContent struct {
	Role  string       `json:"role"`
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiGenerationConfig struct {
	MaxOutputTokens int      `json:"maxOutputTokens,omitempty"`
	Temperature     *float64 `json:"temperature,omitempty"`
	TopP            *float64 `json:"topP,omitempty"`
	StopSequences   []string `json:"stopSequences,omitempty"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
	} `json:"candidates"`
	UsageMetadata struct {
		PromptTokenCount     int `json:"promptTokenCount"`
		CandidatesTokenCount int `json:"candidatesTokenCount"`
		TotalTokenCount      int `json:"totalTokenCount"`
	} `json:"usageMetadata"`
}

func (p *GoogleProvider) ChatCompletion(ctx context.Context, req models.ChatCompletionRequest) (*models.ChatCompletionResponse, error) {
	gemReq := p.toGeminiRequest(req)

	body, err := json.Marshal(gemReq)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/v1beta/models/%s:generateContent?key=%s", p.baseURL, req.Model, p.apiKey)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("google API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var gemResp geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&gemResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(gemResp.Candidates) == 0 || len(gemResp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no content in response")
	}

	content := ""
	for _, part := range gemResp.Candidates[0].Content.Parts {
		content += part.Text
	}

	return &models.ChatCompletionResponse{
		ID:       fmt.Sprintf("gemini-%d", time.Now().UnixNano()),
		Provider: "google",
		Model:    req.Model,
		Message: models.ChatMessage{
			Role:    "assistant",
			Content: content,
		},
		Usage: models.TokenUsage{
			PromptTokens:     gemResp.UsageMetadata.PromptTokenCount,
			CompletionTokens: gemResp.UsageMetadata.CandidatesTokenCount,
			TotalTokens:      gemResp.UsageMetadata.TotalTokenCount,
		},
	}, nil
}

func (p *GoogleProvider) StreamCompletion(ctx context.Context, req models.ChatCompletionRequest, chunks chan<- models.StreamChunk) error {
	defer close(chunks)

	gemReq := p.toGeminiRequest(req)

	body, err := json.Marshal(gemReq)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/v1beta/models/%s:streamGenerateContent?alt=sse&key=%s", p.baseURL, req.Model, p.apiKey)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("google API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")

		var gemResp geminiResponse
		if err := json.Unmarshal([]byte(data), &gemResp); err != nil {
			continue
		}

		if len(gemResp.Candidates) > 0 && len(gemResp.Candidates[0].Content.Parts) > 0 {
			for _, part := range gemResp.Candidates[0].Content.Parts {
				if part.Text != "" {
					chunks <- models.StreamChunk{
						Provider: "google",
						Model:    req.Model,
						Delta:    part.Text,
					}
				}
			}
		}
	}

	chunks <- models.StreamChunk{Provider: "google", Model: req.Model, Done: true}
	return scanner.Err()
}

func (p *GoogleProvider) Models() []models.ModelInfo {
	return []models.ModelInfo{
		{ID: "gemini-2.0-flash", Provider: "google", Name: "Gemini 2.0 Flash", MaxTokens: 1048576, InputPricing: 0.0001, OutputPricing: 0.0004},
		{ID: "gemini-1.5-pro", Provider: "google", Name: "Gemini 1.5 Pro", MaxTokens: 2097152, InputPricing: 0.00125, OutputPricing: 0.005},
		{ID: "gemini-1.5-flash", Provider: "google", Name: "Gemini 1.5 Flash", MaxTokens: 1048576, InputPricing: 0.000075, OutputPricing: 0.0003},
	}
}

func (p *GoogleProvider) toGeminiRequest(req models.ChatCompletionRequest) geminiRequest {
	var contents []geminiContent
	for _, m := range req.Messages {
		role := m.Role
		if role == "assistant" {
			role = "model"
		}
		if role == "system" {
			role = "user"
		}
		contents = append(contents, geminiContent{
			Role:  role,
			Parts: []geminiPart{{Text: m.Content}},
		})
	}

	var genConfig *geminiGenerationConfig
	if req.MaxTokens > 0 || req.Temperature != nil || req.TopP != nil || len(req.Stop) > 0 {
		genConfig = &geminiGenerationConfig{
			MaxOutputTokens: req.MaxTokens,
			Temperature:     req.Temperature,
			TopP:            req.TopP,
			StopSequences:   req.Stop,
		}
	}

	return geminiRequest{
		Contents:         contents,
		GenerationConfig: genConfig,
	}
}
