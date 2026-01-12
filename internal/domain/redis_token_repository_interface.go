package domain

import (
	"context"
	"time"
)

type RedisTokenRepository interface {
	SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error
	DeleteByToken(ctx context.Context, token string) error
	GetUserIDByToken(ctx context.Context, token string) (uint, error)
}