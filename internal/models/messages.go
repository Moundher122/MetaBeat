package models

import (
	"metabeat/pkg/Base"
	
)

type Message struct {
	Base.BaseModel
	Content   string `gorm:"not null"`
	UserID    *uint  `gorm:"foreignKey;null"`
	IPAddress string `gorm:"not null"`
}
