package strategy

import (
	"sync"
	"time"
)

type MemoryStrategy struct {
	mu       sync.Mutex
	limits   map[string]*RateLimit
}

type RateLimit struct {
	Count      int
	ResetAfter time.Time
}

func NewMemoryStrategy() *MemoryStrategy {
	return &MemoryStrategy{
		limits: make(map[string]*RateLimit),
	}
}

func (m *MemoryStrategy) AllowRequest(key string, maxRequests int, duration time.Duration) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	limit, exists := m.limits[key]

	if !exists || now.After(limit.ResetAfter) {
		m.limits[key] = &RateLimit{
			Count:      1,
			ResetAfter: now.Add(duration),
		}
		return true
	}

	if limit.Count < maxRequests {
		limit.Count++
		return true
	}

	return false
}

func (m *MemoryStrategy) Reset(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.limits, key)
}