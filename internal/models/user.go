package models

import (
	"metabeat/pkg/Base"
	"metabeat/pkg/Hash"
)

type User struct {
	Base.BaseModel
	hash.Hash
	Username string `gorm:"unique;not null;primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Bio      string `gorm:"null"`
	BgSong   string `gorm:"null"`
}

