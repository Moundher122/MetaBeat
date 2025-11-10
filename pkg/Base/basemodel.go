package Base

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (obj *BaseModel) SoftDelete() {
	obj.DeletedAt.Time = time.Now()
}
