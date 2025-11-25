package services
import (
	"encoding/json"
	"metabeat/internal/dto"
	"metabeat/internal/models"
	"net/http"
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
	v := h.Validator
	if err := v.Struct(body); err != nil {
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

