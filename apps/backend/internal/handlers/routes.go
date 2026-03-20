package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ssbreno/white-label-repository/backend/internal/config"
	"github.com/ssbreno/white-label-repository/backend/internal/database"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *gin.Engine, db *database.DB, cfg *config.Config) {
	// API group
	api := router.Group("/api")

	// Health check
	api.GET("/health", healthHandler(db, cfg))

	// Users
	userHandler := newUserHandler(db)
	users := api.Group("/users")
	{
		users.GET("", userHandler.list)
		users.GET("/:id", userHandler.getByID)
		users.POST("", userHandler.create)
		users.PUT("/:id", userHandler.update)
		users.DELETE("/:id", userHandler.delete)
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
