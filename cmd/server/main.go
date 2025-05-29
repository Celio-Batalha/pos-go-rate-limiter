package main

import (
	"log"

	"github.com/Celio-Batalha/app-ratelimiter/config"
	"github.com/Celio-Batalha/app-ratelimiter/internal/middleware"
	"github.com/Celio-Batalha/app-ratelimiter/internal/ratelimiter"
	"github.com/Celio-Batalha/app-ratelimiter/internal/ratelimiter/strategy"
	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Cria storage Redis
	storage := strategy.NewRedisStorage(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)

	// Cria rate limiter
	limiter := ratelimiter.NewLimiter(storage, cfg)

	// Cria middleware
	rateLimitMiddleware := middleware.NewRateLimitMiddleware(limiter)

	// Configura Gin
	r := gin.Default()

	// Aplica middleware
	r.Use(rateLimitMiddleware.Handle)

	// Rota de teste
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World!"})
	})

	// Inicia servidor
	log.Printf("Server starting on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
