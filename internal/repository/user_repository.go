package repository

import (
	"gorm.io/gorm"
	"learn_clean_architecture/internal/entity"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}