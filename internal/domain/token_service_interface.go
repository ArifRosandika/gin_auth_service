package domain

import "context"

type TokenService interface {
	GenerateAccessToken(ctx context.Context, userID uint, email string) (string, error)
	GenerateRefreshToken(ctx context.Context, userID uint) (string, error)
	ValidateToken(ctx context.Context, token string) (*RefreshTokenClaims, error)
}