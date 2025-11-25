package endpoint

import (
	v1 "metabeat/internal/endpoint/v1"
	"metabeat/internal/services"
	"net/http"
)

func RegisterEndpoints(mux *http.ServeMux, handler *services.Handler) {
	v1.AuthEndpoints(mux, handler)
	
}
