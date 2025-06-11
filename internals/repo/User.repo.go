package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/khandev-bac/lemon/internals/db/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) FindById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	res := r.db.WithContext(ctx).First(&user, "id = ?", id)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", res.Error)
	}
	return &user, nil
}
func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	res := r.db.WithContext(ctx).First(&user, "email = ?", email)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user: %w", res.Error)
	}
	return &user, nil
}

func (r *UserRepo) UpdateRefreshToken(ctx context.Context, id uuid.UUID, refreshToken string) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("refresh_token", refreshToken).Error
}

func (r *UserRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, "id = ?", id).Error
}

func (r *UserRepo) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("password", password).Error
}
func (r *UserRepo) UpdateEmail(ctx context.Context, id uuid.UUID, email string) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("user_email", email).Error
}

// create user done
// find by id
//find by email
// update password
// update email
// deleteby id
