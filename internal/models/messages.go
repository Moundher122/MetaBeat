package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint   `gorm:"foreignKey;not null"`
}
