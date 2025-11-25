package models

import (
	"metabeat/pkg/Base"
)

type View struct {
	Base.BaseModel
	UserID    *uint  `gorm:"foreignKey;null"`
	Origin    string `gorm:"not null"`
	IPAddress string `gorm:"not null"`
}
