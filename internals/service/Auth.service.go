package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/khandev-bac/lemon/internals/db/model"
	"github.com/khandev-bac/lemon/internals/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewService(repo *repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) Signup(ctx context.Context, user model.User) (*model.User, error) {
	existUser, err := s.repo.FindByEmail(ctx, user.UserEmail)
	if err == nil && existUser != nil {
		return nil, fmt.Errorf("user already exists")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	newUser := &model.User{
		ID:           uuid.New(),
		UserName:     user.UserName,
		UserEmail:    user.UserEmail,
		Password:     string(hashedPass),
		AuthProvider: "auth",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
	if err := s.repo.CreateUser(ctx, newUser); err != nil {
		return nil, fmt.Errorf("failed to create user")
	}
	return newUser, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid email/password")
	}
	return user, nil
}
func (s *UserService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.FindByEmail(ctx, email)
}
func (s *UserService) FindById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return s.repo.FindById(ctx, id)
}
func (s *UserService) UpdateRefreshToken(ctx context.Context, id uuid.UUID, refreshToken string) error {
	return s.repo.UpdateRefreshToken(ctx, id, refreshToken)
}
func (s *UserService) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	return s.repo.UpdatePassword(ctx, id, password)
}

func (s *UserService) UpdateEmail(ctx context.Context, id uuid.UUID, email string) error {
	return s.repo.UpdateEmail(ctx, id, email)
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteUser(ctx, id)
}
