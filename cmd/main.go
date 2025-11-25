package main

import (
	"metabeat/internal/db"
	"metabeat/internal/endpoint"
	"metabeat/internal/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	DB := db.InitDB()
	DB.Migrate()
	mux := http.NewServeMux()
	endpoint.RegisterEndpoints(mux, &services.Handler{DB: DB, Validator: validator.New()})
	http.ListenAndServe(":8080", mux)
}
