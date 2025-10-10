package api

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"` // хранить хэш, не plain!
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RefreshToken struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index"`
	TokenHash string `gorm:"not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &RefreshToken{})
}
