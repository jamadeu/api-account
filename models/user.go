package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Document  string `gorm:"not null,unique"`
	Email     string `gorm:"not null,unique"`
	AccountID uint
}
