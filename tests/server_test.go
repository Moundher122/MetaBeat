package tests

import (
	"metabeat/cmd/server"
	"metabeat/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestHealthEndpoint(t *testing.T) {
	db := db.InitDB()
	db.Migrate()
	go server.Main(db, validator.New())
	req := httptest.NewRequest("GET", "http://localhost:8080/health", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
}
