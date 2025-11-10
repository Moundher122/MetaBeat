package models

import "metabeat/pkg/Base"

type Analytics struct {
	Base.BaseModel
	PageViews    int `gorm:"not null"`
	UniqueVisits int `gorm:"not null"`
}
