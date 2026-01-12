package usecase

import (
	"context"
	"log"
	"errors"
	"fmt"
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"learn_clean_architecture/internal/delivery/http/dto/response"
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/helper"
)

type authUseCaseImpl struct {
	tokenSrv domain.TokenService
	tokenRepo domain.RedisTokenRepository
	userRepo domain.UserRepository
}

func NewAuthUseCase(tokenSrv domain.TokenService, userRepo domain.UserRepository, tokenRepo domain.RedisTokenRepository) domain.AuthUseCase {
	return &authUseCaseImpl{
		tokenSrv: tokenSrv,
		tokenRepo: tokenRepo,
		userRepo: userRepo,
	}
}

func (u *authUseCaseImpl) Login(ctx context.Context, req request.LoginUserRequest) (response.LoginUserResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		return response.LoginUserResponse{}, errors.New("email not found")
	}

	if !helper.CheckPassword(user.Password, req.Password) {
		return response.LoginUserResponse{}, errors.New("wrong password")
	}

	access, err := u.tokenSrv.GenerateAccessToken(ctx, user.ID, user.Email)

	if err != nil {
		return response.LoginUserResponse{}, fmt.Errorf("failed generating token %w", err)
	}

	refresh, err := u.tokenSrv.GenerateRefreshToken(ctx, user.ID)

	if err != nil {
		return response.LoginUserResponse{}, err
	}

	return response.LoginUserResponse{
		AccessToken : access,
		RefreshToken : refresh,
	}, nil
}

func (u *authUseCaseImpl) Refresh(ctx context.Context, refreshToken string) (string, error) {

	userID, err := u.tokenRepo.GetUserIDByToken(ctx, refreshToken)
	if err != nil {
		return "", errors.New("refresh token revoked or invalid")
	}
	fmt.Print("token:", refreshToken)
	fmt.Print("userID:", userID)

	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil || user == nil {
		fmt.Print("userID:", userID)
		return "", errors.New("user not found")
	}

	access, err := u.tokenSrv.GenerateAccessToken(ctx, user.ID, user.Email)

	if err != nil {
		return "", err
	}

	if err := u.tokenRepo.DeleteByToken(ctx, refreshToken); err != nil {
		log.Println(err)
	}

	return access, nil
}