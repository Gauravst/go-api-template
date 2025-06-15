package database

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(uri string) (*gorm.DB, error) {
	if uri == "" {
		return nil, errors.New("URI is empty")
	}
	return gorm.Open(postgres.Open(uri), &gorm.Config{})
}
