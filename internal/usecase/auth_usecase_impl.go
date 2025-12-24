package usecase

import (
	"context"
	"errors"
	"fmt"
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"learn_clean_architecture/internal/delivery/http/dto/response"
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/helper"
)

type authUseCaseImpl struct {
	userRepo domain.UserRepository
	tokenSrv domain.TokenService
	authRepo domain.AuthRepository
}

type RefreshTokenClaims struct {
	UserID uint
	Email string
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

	if err := u.authRepo.SaveRefreshToken(ctx, user.ID, refresh, RefreshTokenTTL); err != nil {
		return response.LoginUserResponse{}, fmt.Errorf("failed saving refresh token %w", err)
	}

	return response.LoginUserResponse{
		AccessToken : access,
		RefreshToken : refresh,
	}, nil
}

func (u *authUseCaseImpl) Refresh(ctx context.Context, refreshToken string) (string, error) {

	claims, err := u.tokenSrv.ValidateToken(ctx, refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	exist, err := u.authRepo.Exists(ctx, claims.UserID, refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed checking refresh token %w", err)
	}

	if !exist {
		return "", errors.New("refresh token revoked")
	}

	access, err := u.tokenSrv.GenerateAccessToken(ctx, claims.UserID, claims.Email)

	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	return access, nil
}