package api

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"` // хранить хэш, не plain!
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
