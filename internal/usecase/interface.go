package usecase

import (
	"learn_clean_architecture/internal/delivery/http/request"
	"learn_clean_architecture/internal/delivery/http/response"
	"learn_clean_architecture/internal/entity"
)

type UserUseCase interface {
	Register(request.RegisterUserRequest) error
	Login(request.LoginUserRequest) (response.LoginUserResponse, error)
	GetProfile(email string) (*entity.User, error)
}