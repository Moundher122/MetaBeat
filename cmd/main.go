package main

import (
	"metabeat/internal/api"
	"metabeat/internal/db"
	"net/http"
)

func main() {
	DB := db.InitDB()
	DB.Migrate()
	http.HandleFunc("/health", api.HealthCheckHandler)
	http.ListenAndServe(":8080", nil)
}
