package services

import (
	"errors"
	"fmt"

	"github.com/gauravst/got/internal/models"
	"github.com/gauravst/got/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("user name cannot be empty")
	}

	err := s.userRepo.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
