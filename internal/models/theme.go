package models

import "metabeat/pkg/Base"

type Theme struct {
	Base.BaseModel
	BgColor     string  `gorm:"not null" default:"#ffffff"`
	BgImage     *string `gorm:"null"`
	TextColor   string  `gorm:"not null" default:"#000000"`
	LinkColor   string  `gorm:"not null" default:"#0000ee"`
	AccentColor string  `gorm:"not null" default:"#6200ee"`
	BorderColor string  `gorm:"not null" default:"#dddddd"`
	UserID      uint    `gorm:"foreignKey;not null"`
}
