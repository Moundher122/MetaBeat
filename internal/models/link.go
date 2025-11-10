package models

import "metabeat/pkg/Base"
type Link struct {
	Base.BaseModel
	URL    string `gorm:"not null"`
	Icon   string `gorm:"null"`
	UserID uint   `gorm:"foreignKey;not null"`
}
