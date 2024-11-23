package config

import (
	"github.com/jamadeu/api-account/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	// Connect DB
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// Migrate the schema
	if err = db.AutoMigrate(&models.Account{}); err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return db, nil
}
