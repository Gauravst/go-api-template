package repositories

import (
	"database/sql"

	"github.com/gauravst/got/internal/models"
)

// AuthRepository defines the interface for user-related database operations
type AuthRepository interface {
	LoginUser(user *models.User) error
}

// authRepository implements the AuthRepository interface
type authRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of userRepository
func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

// CreateUser inserts a new user into the database
func (r *authRepository) LoginUser(user *models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}
