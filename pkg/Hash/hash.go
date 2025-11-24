package hash

import (
	"golang.org/x/crypto/bcrypt"
)
type Hash struct{}

func (h *Hash) HashWord(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func (h *Hash) CheckWord(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
