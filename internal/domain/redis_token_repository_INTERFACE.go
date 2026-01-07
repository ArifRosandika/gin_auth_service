package domain

import (
	"context"
	"time"
)

type RedisTokenRepository interface {
	SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error
	GetRefreshToken(ctx context.Context, userID uint) (string, error)
	BlackListToken(ctx context.Context, token string, ttl time.Duration) error
	IsBlacklisted(ctx context.Context, token string) (bool, error)
	Exists(ctx context.Context, userID uint, token string) (bool, error)
}