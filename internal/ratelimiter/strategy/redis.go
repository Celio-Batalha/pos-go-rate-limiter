package strategy

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr, password string, db int) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisStorage) Get(key string) (int, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	count, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *RedisStorage) Set(key string, value int, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisStorage) Increment(key string, expiration time.Duration) (int, error) {
	pipe := r.client.TxPipeline()

	incr := pipe.Incr(r.ctx, key)
	pipe.Expire(r.ctx, key, expiration)

	_, err := pipe.Exec(r.ctx)
	if err != nil {
		return 0, err
	}

	return int(incr.Val()), nil
}

func (r *RedisStorage) IsBlocked(key string) (bool, error) {
	blockedKey := fmt.Sprintf("blocked:%s", key)
	exists, err := r.client.Exists(r.ctx, blockedKey).Result()
	return exists > 0, err
}

func (r *RedisStorage) Block(key string, duration time.Duration) error {
	blockedKey := fmt.Sprintf("blocked:%s", key)
	return r.client.Set(r.ctx, blockedKey, 10, duration).Err()
}
