package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite(dbPath string) (*gorm.DB, error) {
	if dbPath == "" {
		return nil, errors.New("db Path is empty")
	}

	return gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
}
