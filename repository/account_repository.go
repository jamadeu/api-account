package repository

import (
	"github.com/jamadeu/api-account/models"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) FindById(id string) (*models.Account, error) {
	a := models.Account{}
	if err := ar.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}
