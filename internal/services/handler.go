package services
import (
	"metabeat/internal/db"
	"github.com/go-playground/validator/v10"
)
type Handler struct {
	DB *db.DB
	Validator *validator.Validate
}