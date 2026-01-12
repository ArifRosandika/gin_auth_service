package usecase

import (
	"context"
	"errors"
	"learn_clean_architecture/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenTTL = 15 *time.Minute
	RefreshTokenTTL = 24 * 7 *time.Hour
)

type jwtClaims struct {
	UserID uint `json:"user_id"`
	Email string `json:"email,omitempty"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secret string
	tokenRepo domain.RedisTokenRepository
}

func NewTokenService(secret string, tokenRepo domain.RedisTokenRepository) domain.TokenService {
	return &jwtService{
		secret: secret,
		tokenRepo: tokenRepo,
	}
}

func (s *jwtService) GenerateAccessToken(ctx context.Context, userID uint, email string) (string, error) {
	claims := jwtClaims{
		UserID : userID,
		Email : email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTTL)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(s.secret))
}

func (s *jwtService) GenerateRefreshToken(ctx context.Context, userID uint) (string, error) {
	claims := jwtClaims{
		UserID : userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenTTL)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := t.SignedString([]byte(s.secret))

	if err != nil {
		return "", err
	}

	if err := s.tokenRepo.SaveRefreshToken(ctx, userID, tokenStr, RefreshTokenTTL); err != nil {
		return "", err
	}
	
	return tokenStr, nil
}

func (s *jwtService) ValidateAccessToken(ctx context.Context, tokenStr string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims := token.Claims.(*jwtClaims)

	return claims.UserID, nil
}

func (s *jwtService) ValidateRefreshToken(ctx context.Context, tokenStr string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims := token.Claims.(*jwtClaims)

	return claims.UserID, nil
}
