package repository

import (
	"context"
	"learn_clean_architecture/internal/domain"
	"time"

	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	DB *gorm.DB
}
func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepositoryImpl{DB: db}
}

func (r *authRepositoryImpl) SaveRefreshToken(ctx context.Context, userID uint, token string, ttl time.Duration) error {
	return r.DB.WithContext(ctx).Create(&domain.RefreshToken{
		UserID: userID,
		Token: token,
		ExpiresAt: time.Now().Add(ttl),
	}).Error
}

func (r *authRepositoryImpl) Exists(ctx context.Context, userID uint, token string) (bool, error) {
	var count int64
	err :=  r.DB.WithContext(ctx).Where("user_id = ? AND token = ?", userID, token).Find(&domain.RefreshToken{}).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *authRepositoryImpl) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.DB.WithContext(ctx).Where("user_id = ?", userID).Delete(&domain.RefreshToken{}).Error
}