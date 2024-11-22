package config

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *zap.Logger
)

func Init() error {
	var err error

	// Connect to database
	db, err = ConnectDb()
	if err != nil {
		return fmt.Errorf("Error connecting db: %v", err)
	}
	return nil
}

func GetConnDB() *gorm.DB {
	return db
}

func GetLogger() *zap.Logger {
	return Logger
}
