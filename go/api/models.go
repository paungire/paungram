package api

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	AvatarURL string `gorm:"default:'/default_avatar.png'"`
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	SentMessages       []Message    `gorm:"foreignKey:SenderID"`
	ReceivedMessages   []Message    `gorm:"foreignKey:ReceiverID"`
	Friendships        []Friendship `gorm:"foreignKey:UserID"`
	InverseFriendships []Friendship `gorm:"foreignKey:FriendID"`
}

type Friendship struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	FriendID  uint   `gorm:"not null"`
	Status    string `gorm:"type:varchar(20);default:'pending'"` // e.g., 'pending', 'accepted', 'blocked'
	CreatedAt time.Time
	UpdatedAt time.Time

	User   User `gorm:"foreignKey:UserID"`
	Friend User `gorm:"foreignKey:FriendID"`
}

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	SenderID   uint      `gorm:"not null"`
	ReceiverID uint      `gorm:"not null"`
	Content    string    `gorm:"type:text;not null"`
	SentAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ReadAt     *time.Time

	Sender   User `gorm:"foreignKey:SenderID"`
	Receiver User `gorm:"foreignKey:ReceiverID"`
}

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	TokenHash string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &RefreshToken{})
}
