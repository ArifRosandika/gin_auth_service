package usecase

import (
	"context"
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/helper"
)

type userUseCaseImpl struct {
	userRepo domain.UserRepository
	tokenRepo domain.RedisTokenRepository
	tokenSrv domain.TokenService
}

func NewUserUseCase(userRepo domain.UserRepository, tokenRepo domain.RedisTokenRepository, tokenSrv domain.TokenService) domain.UserUseCase {
	return &userUseCaseImpl{
		userRepo: userRepo,
		tokenRepo: tokenRepo,
		tokenSrv: tokenSrv,
	}
}

func (u *userUseCaseImpl) Register(ctx context.Context, req request.RegisterUserRequest) error {
	hashed, _ := helper.HashPassword(req.Password)
	user := &domain.User{
		Name: req.Name,
		Email: req.Email,
		Password: hashed,
	}
	return u.userRepo.Create(ctx, user)
}

func (u *userUseCaseImpl) GetProfile(ctx context.Context, email string) (*domain.User, error) {
	return u.userRepo.FindByEmail(ctx, email)
}
