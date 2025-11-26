package server
import (
	"fmt"
	"metabeat/internal/db"
	"metabeat/internal/endpoint"
	"metabeat/internal/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func Main(db *db.DB, validator *validator.Validate) error {
	db.Migrate()
	mux := http.NewServeMux()
	endpoint.RegisterEndpoints(mux, &services.Handler{DB: db, Validator: validator})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}
	return nil
}