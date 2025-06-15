package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/gauravst/got/internal/models"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbType string
)

// createConnection initializes a single DB connection
func CreateConnection(driver string, url string) (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	switch driver {
	case "postgres":
		db, err = InitPostgres(url)
	case "sqlite":
		db, err = InitSQLite(url)
	default:
		return nil, fmt.Errorf("unsupported db type: %s", driver)
	}

	if err != nil {
		return nil, err
	}

	// Auto-migrate schema
	err = db.AutoMigrate(
		&models.User{},
		// add more models here
	)
	if err != nil {
		return nil, fmt.Errorf("auto-migrate failed: %v", err)
	}
	log.Println("Auto-migrated all schemas")

	dbType = driver
	log.Printf("DB connected [%s]", dbType)
	return db, nil
}

// GetConnection returns the active DB instance
func GetConnection() (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("DB connection not initialized")
	}
	return db, nil
}

// CloseConnection closes the active DB connection
func CloseConnection() error {
	if db == nil {
		return errors.New("no DB connection to close")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("error getting sql.DB: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("error closing DB: %v", err)
	}

	db = nil
	log.Println("DB connection closed")
	return nil
}
