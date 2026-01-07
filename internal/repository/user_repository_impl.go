package repository

import (
	"context"
	"learn_clean_architecture/internal/domain"
	"time"

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
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error {
	return r.DB.Model(&domain.RefreshToken{}).Where("refresh_token = ?", token).Update("refresh_token", token).Error
}

func (r *userRepositoryImpl) Exists(ctx context.Context, userID uint, token string) (bool, error) {
	var user domain.User

	return r.DB.Where("id = ? AND refresh_token", userID, token).First(&user).RowsAffected > 0, nil	
}


func (r *userRepositoryImpl) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.DB.Where("id = ?", userID).Delete(&domain.User{}).Error
}