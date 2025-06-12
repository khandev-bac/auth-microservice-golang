package dto

import (
	"time"
)

type SignupDTO struct {
	UserName    string    `json:"user_name" validate:"required,min=3,max=100"`
	UserEmail   string    `json:"user_email" validate:"required,email"`
	Password    string    `json:"password" validate:"required,min=6"`
	UserProfile string    `json:"user_profile" validate:"omitempty,url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type LoginDTO struct {
	UserEmail string `json:"user_email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=100"`
}
