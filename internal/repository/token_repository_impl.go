package repository

import (
	"errors"
	"fmt"
	"learn_clean_architecture/internal/cache"
	"learn_clean_architecture/internal/domain"
	"strconv"
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
	key := fmt.Sprintf("refresh:%s", token)
	val := strconv.Itoa(int(userID))

	return r.cache.Set(ctx, key, val, ttl)
}

func (r *redisTokenRepo) GetUserIDByToken(ctx context.Context, token string) (uint, error) {
	key := fmt.Sprintf("refresh:%s", token)
	val, err := r.cache.Get(ctx, key)

	if err != nil {
		return 0, err
	}

	
	id,err := strconv.ParseUint(val, 10, 64)
	if val == "" {
		return 0, errors.New("token revoked or invalid")
	}
	
	return uint(id), nil
}

func (r *redisTokenRepo) DeleteByToken(ctx context.Context, token string) error {
	key := fmt.Sprintf("refresh:%s", token)
	return r.cache.Delete(ctx, key)
}