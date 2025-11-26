package services

import (
	"encoding/json"
	"fmt"
	"metabeat/internal/dto"
	"metabeat/internal/models"
	"net/http"

	"gorm.io/gorm"
)

func (h *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var body dto.RegisterDto
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Validator.Struct(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := models.User{
		Email:    body.Email,
		Password: body.Password,
		Username: body.Username,
	}
	user.Password, _ = user.HashWord(user.Password)
	if err := h.DB.Db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/js on")
	w.WriteHeader(http.StatusCreated)
	response := dto.UserDto{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		BgSong:   user.BgSong,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var body dto.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	if body.Email != nil {
		user := gorm.G[models.User](h.DB.Db).Where("email = ?", *body.Email)
		fmt.Println(user)
	} else {

	}
	fmt.Println(user.Password, body.Password)
	if !user.CheckWord(body.Password, user.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := dto.UserDto{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		BgSong:   user.BgSong,
	}
	json.NewEncoder(w).Encode(response)

}
func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {}
