package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	URL    string `gorm:"not null"`
	Icon   string `gorm:"null"`
	UserID uint   `gorm:"foreignKey;not null"`
}
