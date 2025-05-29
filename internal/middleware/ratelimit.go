package middleware

import (
	"net/http"

	"github.com/Celio-Batalha/app-ratelimiter/internal/ratelimiter"
	"github.com/gin-gonic/gin"
)

type RateLimitMiddleware struct {
	limiter *ratelimiter.Limiter
}

func NewRateLimitMiddleware(limiter *ratelimiter.Limiter) *RateLimitMiddleware {
	return &RateLimitMiddleware{limiter: limiter}
}

func (m *RateLimitMiddleware) Handle(c *gin.Context) {
	ip := c.ClientIP()
	token := c.GetHeader("API_KEY")

	if m.limiter.Exceeded(ip, token) {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"Código Http": http.StatusTooManyRequests,
			"message":     "you have reached the maximum number of requests or actions allowed within a certain time frame",
		})
		c.Abort()
		return
	}

	m.limiter.Increment(ip, token)
	c.Next()
}
