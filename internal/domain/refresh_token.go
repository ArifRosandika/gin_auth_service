package domain

import "time"

type RefreshToken struct {
	ID uint `gorm:"primary_key"`
	UserID uint `gorm:"not null"`
	Token string `gorm:"not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
}