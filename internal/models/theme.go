package models

import "metabeat/pkg/Base"

type Theme struct {
	Base.BaseModel
	BgColor     string `gorm:"not null"`
	TextColor   string `gorm:"not null"`
	LinkColor   string `gorm:"not null"`
	AccentColor string `gorm:"not null"`
	BorderColor string `gorm:"not null"`
	UserID      uint   `gorm:"foreignKey;not null"`
}
