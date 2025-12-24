package domain

import (
	"context"
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"learn_clean_architecture/internal/delivery/http/dto/response"
)

type AuthUseCase interface {
	Login (ctx context.Context, req request.LoginUserRequest) (response.LoginUserResponse, error)
	Refresh (ctx context.Context, refreshToken string) (string, error)
}