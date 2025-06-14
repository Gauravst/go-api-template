package services

import (
	"errors"
	"fmt"

	"github.com/gauravst/got/internal/models"
	"github.com/gauravst/got/internal/repositories"
)

type AuthService interface {
	LoginUser(user *models.User) error
}

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return &authService{
		authRepo: authRepo,
	}
}

func (s *authService) LoginUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("user name cannot be empty")
	}

	err := s.authRepo.LoginUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
