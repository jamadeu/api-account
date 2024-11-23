package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	User    User
	Balance float64 `gorm:"not null"`
}

type AccountRepository interface {
	FindById(id string) (*Account, error)
}
