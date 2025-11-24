package Base

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (obj *BaseModel) SoftDelete() {
	obj.DeletedAt.Time = time.Now()
}
