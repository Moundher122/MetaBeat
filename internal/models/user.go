package models

import (
	"metabeat/pkg/Base"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base.BaseModel
	Username string `gorm:"unique;not null;primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Bio      string `gorm:"null"`
	BgSong   string `gorm:"null"`
}

func (u *User) HashPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)
}
