package models

import "metabeat/pkg/Base"
type Media struct {
	Base.BaseModel
	FileUrl string `gorm:"not null"`
	Size    string `gorm:"not null"`
}
