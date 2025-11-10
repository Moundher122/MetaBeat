package models

import (
	"metabeat/pkg/Base"
)

type User struct {
	Base.BaseModel
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Bio      string `gorm:"null"`
	BgSong   string `gorm:"null"`
}
