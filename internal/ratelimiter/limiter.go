package ratelimiter

import (
	"fmt"
	"time"

	"github.com/Celio-Batalha/app-ratelimiter/config"
	"github.com/Celio-Batalha/app-ratelimiter/internal/ratelimiter/strategy"
)

type Limiter struct {
	storage strategy.Storage
	config  *config.Config
}

func NewLimiter(storage strategy.Storage, config *config.Config) *Limiter {
	return &Limiter{
		storage: storage,
		config:  config,
	}
}

func (l *Limiter) Exceeded(ip, token string) bool {
	// Prioridade: Token > IP
	if token != "" {
		return l.checkLimit(fmt.Sprintf("token:%s", token), l.config.TokenLimit, l.config.TokenBlockTime)
	}

	return l.checkLimit(fmt.Sprintf("ip:%s", ip), l.config.IPLimit, l.config.IPBlockTime)
}

func (l *Limiter) Increment(ip, token string) {
	if token != "" {
		l.incrementCounter(fmt.Sprintf("token:%s", token), l.config.TokenLimit, l.config.TokenBlockTime)
	} else {
		l.incrementCounter(fmt.Sprintf("ip:%s", ip), l.config.IPLimit, l.config.IPBlockTime)
	}
}

func (l *Limiter) checkLimit(key string, limit int, blockTime time.Duration) bool {
	// Verifica se está bloqueado
	blocked, err := l.storage.IsBlocked(key)
	if err != nil {
		return false
	}
	if blocked {
		return true
	}

	// Verifica o contador atual
	currentCount, err := l.storage.Get(key)
	if err != nil {
		return false
	}

	return currentCount >= limit
}

func (l *Limiter) incrementCounter(key string, limit int, blockTime time.Duration) {
	// Incrementa o contador
	count, err := l.storage.Increment(key, time.Second)
	if err != nil {
		return
	}

	// Se excedeu o limite, bloqueia
	if count > limit {
		l.storage.Block(key, blockTime)
	}
}
