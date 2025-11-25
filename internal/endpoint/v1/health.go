package v1

import (
	"metabeat/internal/services"
	"net/http"
)

func HealthEndpoints(mux *http.ServeMux, handler *services.Handler) {
	mux.HandleFunc("/health", handler.HealthCheckHandler)
}
