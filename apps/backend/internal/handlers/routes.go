package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ssbreno/white-label-repository/backend/internal/ai"
	"github.com/ssbreno/white-label-repository/backend/internal/config"
	"github.com/ssbreno/white-label-repository/backend/internal/database"
	"github.com/ssbreno/white-label-repository/backend/internal/middleware"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
	"github.com/ssbreno/white-label-repository/backend/internal/repositories"
	"github.com/ssbreno/white-label-repository/backend/internal/services"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine, db *database.DB, cfg *config.Config) {
	// API group
	api := router.Group("/api")

	// Health check
	api.GET("/health", healthHandler(db, cfg))

	// Users
	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userH := newUserHandler(userSvc)
	users := api.Group("/users")
	{
		users.GET("", userH.list)
		users.GET("/:id", userH.getByID)
		users.POST("", userH.create)
		users.PUT("/:id", userH.update)
		users.DELETE("/:id", userH.delete)
	}

	// AI Gateway
	aiRegistry := ai.NewRegistry(cfg.AI)
	aiH := newAIHandler(aiRegistry)
	aiGroup := api.Group("/ai")
	aiGroup.Use(middleware.RateLimiter(cfg.AI.RateLimit.RPS, cfg.AI.RateLimit.Burst))
	{
		aiGroup.POST("/chat", aiH.chat)
		aiGroup.GET("/models", aiH.listModels)
		aiGroup.GET("/providers", aiH.listProviders)
	}
}

// healthHandler returns a health check endpoint handler
func healthHandler(db *database.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		services := map[string]string{
			"api": "healthy",
		}

		// Check database connectivity
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		dbStatus := "healthy"
		if err := db.Ping(ctx); err != nil {
			dbStatus = "unhealthy: " + err.Error()
		}
		services["database"] = dbStatus

		status := "healthy"
		httpStatus := http.StatusOK
		if dbStatus != "healthy" {
			status = "degraded"
			httpStatus = http.StatusServiceUnavailable
		}

		c.JSON(httpStatus, models.HealthResponse{
			Status:    status,
			Timestamp: time.Now().UTC(),
			Version:   "1.0.0",
			Services:  services,
		})
	}
}
