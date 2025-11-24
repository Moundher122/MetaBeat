package main

import (
	"metabeat/internal/api"
	"metabeat/internal/db"
	"net/http"
)

func main() {
	DB := db.InitDB()
	DB.Migrate()
	Handler := &api.Handler{DB: DB}
	http.HandleFunc("/health", Handler.HealthCheckHandler)
	http.HandleFunc("/register", Handler.RegisterHandler)

	http.ListenAndServe(":8080", nil)
}
