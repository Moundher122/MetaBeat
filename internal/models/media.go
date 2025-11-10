package models

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	FileUrl string `gorm:"not null"`
	Size    string `gorm:"not null"`
}
