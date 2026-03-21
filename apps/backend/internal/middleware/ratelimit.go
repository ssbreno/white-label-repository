package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ssbreno/white-label-repository/backend/internal/models"
)

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	rps      int
	burst    int
}

type visitor struct {
	tokens    float64
	lastCheck time.Time
}

// RateLimiter returns a Gin middleware that limits requests per IP using a token bucket algorithm
func RateLimiter(rps int, burst int) gin.HandlerFunc {
	if rps <= 0 {
		rps = 10
	}
	if burst <= 0 {
		burst = 20
	}

	rl := &rateLimiter{
		visitors: make(map[string]*visitor),
		rps:      rps,
		burst:    burst,
	}

	// Cleanup stale entries every 5 minutes
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			rl.cleanup()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !rl.allow(ip) {
			c.JSON(http.StatusTooManyRequests, models.NewError[any]("rate limit exceeded"))
			c.Abort()
			return
		}
		c.Next()
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	now := time.Now()

	if !exists {
		rl.visitors[ip] = &visitor{
			tokens:    float64(rl.burst) - 1,
			lastCheck: now,
		}
		return true
	}

	elapsed := now.Sub(v.lastCheck).Seconds()
	v.tokens += elapsed * float64(rl.rps)
	if v.tokens > float64(rl.burst) {
		v.tokens = float64(rl.burst)
	}
	v.lastCheck = now

	if v.tokens < 1 {
		return false
	}

	v.tokens--
	return true
}

func (rl *rateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	cutoff := time.Now().Add(-10 * time.Minute)
	for ip, v := range rl.visitors {
		if v.lastCheck.Before(cutoff) {
			delete(rl.visitors, ip)
		}
	}
}
