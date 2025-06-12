package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserName      string    `gorm:"type:text;not null;index"`
	UserEmail     string    `gorm:"type:text;unique;not null"`
	Password      string    `gorm:"type:text;not null"`
	UserProfile   string    `gorm:"type:text"`
	FirebaseId    string    `gorm:"type:text"`
	AuthProvider  string    `gorm:"type:text;default:password"`
	EmailVerified bool      `gorm:"default:false"`
	RefreshToken  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
