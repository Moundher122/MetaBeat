package models
import "gorm.io/gorm"

type Analytics struct {
	gorm.Model
	PageViews   int `gorm:"not null"`
	UniqueVisits int `gorm:"not null"`
}

