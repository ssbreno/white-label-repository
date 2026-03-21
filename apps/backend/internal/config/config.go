package config

import (
	"os"
	"strconv"
	"strings"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	CORS     CORSConfig
	JWT      JWTConfig
	AI       AIConfig
}

// AIConfig holds AI provider configuration
type AIConfig struct {
	OpenAI    ProviderConfig
	Anthropic ProviderConfig
	Google    ProviderConfig
	RateLimit RateLimitConfig
}

// ProviderConfig holds configuration for a single AI provider
type ProviderConfig struct {
	APIKey  string
	BaseURL string
	Enabled bool
}

// RateLimitConfig holds rate limiting configuration
type RateLimitConfig struct {
	RPS   int
	Burst int
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port string
	Env  string
}

// DatabaseConfig holds PostgreSQL connection configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string
}

// JWTConfig holds JWT authentication configuration
type JWTConfig struct {
	Secret     string
	Expiration string
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("BACKEND_PORT", "8080"),
			Env:  getEnv("APP_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "whitelabel_db"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		CORS: CORSConfig{
			AllowedOrigins: strings.Split(
				getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
				",",
			),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "change-this-secret"),
			Expiration: getEnv("JWT_EXPIRATION", "24h"),
		},
		AI: AIConfig{
			OpenAI: ProviderConfig{
				APIKey:  getEnv("OPENAI_API_KEY", ""),
				BaseURL: getEnv("OPENAI_BASE_URL", "https://api.openai.com/v1"),
				Enabled: getEnv("OPENAI_API_KEY", "") != "",
			},
			Anthropic: ProviderConfig{
				APIKey:  getEnv("ANTHROPIC_API_KEY", ""),
				BaseURL: getEnv("ANTHROPIC_BASE_URL", "https://api.anthropic.com"),
				Enabled: getEnv("ANTHROPIC_API_KEY", "") != "",
			},
			Google: ProviderConfig{
				APIKey:  getEnv("GOOGLE_API_KEY", ""),
				BaseURL: getEnv("GOOGLE_BASE_URL", "https://generativelanguage.googleapis.com"),
				Enabled: getEnv("GOOGLE_API_KEY", "") != "",
			},
			RateLimit: RateLimitConfig{
				RPS:   getEnvInt("RATE_LIMIT_RPS", 10),
				Burst: getEnvInt("RATE_LIMIT_BURST", 20),
			},
		},
	}
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
