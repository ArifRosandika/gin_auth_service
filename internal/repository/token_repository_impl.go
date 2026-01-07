package repository

import (
	"fmt"
	"learn_clean_architecture/internal/cache"
	"learn_clean_architecture/internal/domain"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)
type redisTokenRepo struct {
	cache *cache.RedisCache
}

func NewRedisTokenRepository(client *redis.Client) domain.RedisTokenRepository {
	return &redisTokenRepo{
		cache: cache.NewRedisCache(client),
	}
}

func (r *redisTokenRepo) SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error {
	key := fmt.Sprintf("refresh: %d", userID)
	return r.cache.Set(ctx, key, token, ttl)
}

func (r *redisTokenRepo) GetRefreshToken(ctx context.Context, userID uint) (string, error) {
	key := fmt.Sprintf("refresh: %d", userID)
	return r.cache.Get(ctx, key)
}

func (r *redisTokenRepo) BlackListToken(ctx context.Context, token string, ttl time.Duration) error {
	key := fmt.Sprintf("blacklist: %s", token)
	return r.cache.Set(ctx, key, "1", ttl)
}

func (r *redisTokenRepo) IsBlacklisted(ctx context.Context, token string) (bool, error) {
	key := fmt.Sprintf("blacklist: %s", token)
	_, err := r.cache.Exists(ctx, key)

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *redisTokenRepo) Exists(ctx context.Context, userID uint, token string) (bool, error) {
	key := fmt.Sprintf("refresh: %d", userID)
	tkn, err := r.cache.Get(ctx, key)

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		return false, err
	}

	return tkn == token, nil
}