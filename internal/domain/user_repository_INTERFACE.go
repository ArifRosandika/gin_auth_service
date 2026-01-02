package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error
	Exists(ctx context.Context, userID uint, token string) (bool, error)
	FindByID(ctx context.Context, userID uint) (*User, error)
	DeleteByUserID(ctx context.Context, userID uint) error
}