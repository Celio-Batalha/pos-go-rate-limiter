package strategy

import "time"

type Storage interface {
	Get(key string) (int, error)
	Set(key string, value int, expiration time.Duration) error
	Increment(key string, expiration time.Duration) (int, error)
	IsBlocked(key string) (bool, error)
	Block(key string, duration time.Duration) error
}
