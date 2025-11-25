package v1

import (
	"metabeat/internal/services"
	"net/http"
)

func AuthEndpoints(mux *http.ServeMux, handler *services.Handler) {
	mux.HandleFunc("/register", handler.RegisterHandler)
	mux.HandleFunc("/login", handler.LoginHandler)
}
