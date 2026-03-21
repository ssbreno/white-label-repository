package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssbreno/white-label-repository/backend/internal/ai"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

type aiHandler struct {
	registry *ai.Registry
}

func newAIHandler(registry *ai.Registry) *aiHandler {
	return &aiHandler{registry: registry}
}

// chat handles POST /api/ai/chat
func (h *aiHandler) chat(c *gin.Context) {
	var req models.ChatCompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.NewError[any](err.Error()))
		return
	}

	provider, err := h.registry.Get(req.Provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewError[any](err.Error()))
		return
	}

	if req.Stream {
		h.streamChat(c, provider, req)
		return
	}

	resp, err := provider.ChatCompletion(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewError[any](err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.NewSuccess(resp, "Chat completion successful"))
}

// streamChat handles streaming responses via SSE
func (h *aiHandler) streamChat(c *gin.Context, provider ai.Provider, req models.ChatCompletionRequest) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	chunks := make(chan models.StreamChunk, 10)

	go func() {
		if err := provider.StreamCompletion(c.Request.Context(), req, chunks); err != nil {
			chunks <- models.StreamChunk{
				Provider: req.Provider,
				Model:    req.Model,
				Delta:    fmt.Sprintf("[error: %s]", err.Error()),
				Done:     true,
			}
		}
	}()

	c.Stream(func(w io.Writer) bool {
		chunk, ok := <-chunks
		if !ok {
			return false
		}
		data, _ := json.Marshal(chunk)
		fmt.Fprintf(w, "data: %s\n\n", data)
		return !chunk.Done
	})
}

// listModels handles GET /api/ai/models
func (h *aiHandler) listModels(c *gin.Context) {
	allModels := h.registry.AllModels()
	if allModels == nil {
		allModels = []models.ModelInfo{}
	}
	c.JSON(http.StatusOK, models.NewSuccess(allModels, "Models retrieved"))
}

// listProviders handles GET /api/ai/providers
func (h *aiHandler) listProviders(c *gin.Context) {
	providers := h.registry.ProviderNames()
	c.JSON(http.StatusOK, models.NewSuccess(providers, "Providers retrieved"))
}
