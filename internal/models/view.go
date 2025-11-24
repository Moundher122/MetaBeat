package models

import (
	"metabeat/pkg/Base"
)

type View struct {
	Base.BaseModel
	UserID    User   `gorm:"foreignKey;not null"`
	Origin    string `gorm:"not null"`
	IPAddress string `gorm:"not null"`
}
