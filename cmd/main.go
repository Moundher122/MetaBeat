package main

import (
	"fmt"
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
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
