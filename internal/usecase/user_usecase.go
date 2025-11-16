package usecase

import (
	"errors"
	"learn_clean_architecture/internal/entity"
	"learn_clean_architecture/internal/helper"
	"learn_clean_architecture/internal/delivery/http/request"
	"learn_clean_architecture/internal/repository"
	"learn_clean_architecture/internal/delivery/http/response"
)

type userUseCaseImpl struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{
		repo: repo,
	}
}

func (u *userUseCaseImpl) Register(req request.RegisterUserRequest) error {
	hashed, _ := helper.HashPassword(req.Password)
	user := &entity.User{
		Name: req.Name,
		Email: req.Email,
		Password: hashed,
	}
	return u.repo.Create(user)
}

func (u *userUseCaseImpl) Login(req request.LoginUserRequest) (response.LoginUserResponse, error) {
	user, err := u.repo.FindByEmail(req.Email)

	if err != nil {
		return response.LoginUserResponse{}, errors.New("email not found")
	}

	if !helper.CheckPassword(user.Password, req.Password) {
		return response.LoginUserResponse{}, errors.New("wrong password")
	}

	token, err := helper.GenerateToken(user.Email)

	if err != nil {
		return response.LoginUserResponse{}, errors.New("failed generating token")
	}

	return response.LoginUserResponse{
		Token: token,
	}, nil
}

func (u *userUseCaseImpl) GetProfile(email string) (*entity.User, error) {
	return u.repo.FindByEmail(email)
}