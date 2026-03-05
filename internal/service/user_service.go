package service

import (
	"context"
	"errors"
	"task-management/internal/model"
	"task-management/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserByNationalID(ctx context.Context, nationalID string) (*model.User, error) {
	return s.Repo.GetByNationalID(ctx, nationalID)
}

func (s *UserService) CreateUser(ctx context.Context, u *model.User) error {
	// Check for duplicate national ID
	existingUser, _ := s.Repo.GetByNationalID(ctx, u.NationalID)
	if existingUser != nil {
		return errors.New("user with this national ID already exists")
	}

	// Basic validation
	if u.FullName == "" || u.Phone == "" || u.NationalID == "" {
		return errors.New("full name, phone, and national ID are required")
	}

	// Call repository to insert
	return s.Repo.Create(ctx, u)
}