package domain

import (
	"context"
	"time"
)

type AuthRepository interface {
	SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error
	Exists(ctx context.Context, userID uint, token string) (bool, error)
	DeleteByUserID(ctx context.Context, userID uint) error
}