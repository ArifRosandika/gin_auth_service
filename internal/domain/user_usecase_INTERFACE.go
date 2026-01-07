package domain

import (
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"context"
)

type UserUseCase interface {
	Register(ctx context.Context, req request.RegisterUserRequest) error
	GetProfile(ctx context.Context, UserID uint) (*User, error)
}