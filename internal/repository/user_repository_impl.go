package repository

import (
	"context"
	"fmt"
	"learn_clean_architecture/internal/domain"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, userID uint) (*domain.User, error) {
	var user domain.User

	if err := r.DB.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		fmt.Print("userID:", userID)
		return nil, err
	}

	return &user, nil
}