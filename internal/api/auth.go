package api
import (
	"gorm.io/gorm"
	"metabeat/internal/models"
	"net/http"
	"encoding/json"
)

func RegisterHandler(DB *gorm.DB, w http.ResponseWriter, r *http.Request) (models.User, error) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return models.User{}, err
	}
	if err := DB.Create(&user).Error; err != nil {
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
	result := DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
