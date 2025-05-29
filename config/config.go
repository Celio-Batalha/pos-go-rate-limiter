package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       int

	IPLimit     int
	IPBlockTime time.Duration

	TokenLimit     int
	TokenBlockTime time.Duration

	ServerPort string
}

func Load() (*Config, error) {
	// Carrega o arquivo .env se existir
	godotenv.Load()

	cfg := &Config{
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),

		IPLimit:     getEnvAsInt("IP_LIMIT", 5),
		IPBlockTime: time.Duration(getEnvAsInt("IP_BLOCK_TIME", 5)) * time.Minute,

		TokenLimit:     getEnvAsInt("TOKEN_LIMIT", 10),
		TokenBlockTime: time.Duration(getEnvAsInt("TOKEN_BLOCK_TIME", 10)) * time.Minute,
		// IPLimit:     getEnvAsInt("IP_LIMIT", 5),
		// IPBlockTime: parseDuration("IP_BLOCK_TIME", 10), // 10 minutos padrão

		// TokenLimit:     getEnvAsInt("TOKEN_LIMIT", 10),
		// TokenBlockTime: parseDuration("TOKEN_BLOCK_TIME", 30), // 30 minutos padrão

		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// func parseDuration(key, defaultValue string) time.Duration {
// 	value := getEnv(key, defaultValue)
// 	duration, err := time.ParseDuration(value)
// 	if err != nil {
// 		// Fallback: tenta como segundos
// 		seconds := getEnvAsInt(key+"_SECONDS", 300)
// 		return time.Duration(seconds) * time.Second
// 	}
// 	return duration
// }
