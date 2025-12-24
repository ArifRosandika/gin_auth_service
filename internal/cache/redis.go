package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		Client: client,
	}
}

func (r *RedisCache) Set(
	ctx context.Context,
	key string,
	value string,
	ttl time.Duration,
) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCache) Get(
	ctx context.Context,
	key string,
) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisCache) Delete(
	ctx context.Context,
	key string,
) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisCache) Exists(
	ctx context.Context, 
	key string,
) (bool, error) {
	res, err := r.Client.Exists(ctx, key).Result()
	return res > 0, err
}