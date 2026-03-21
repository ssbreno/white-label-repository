package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// NewTestRouter creates a Gin engine in test mode
func NewTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// PerformRequest sends an HTTP request to the router and returns the response recorder
func PerformRequest(router *gin.Engine, method, path string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	return w
}
