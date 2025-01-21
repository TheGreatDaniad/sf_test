package api

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRoutes() *Routes {
	return &Routes{
		SequenceHandler: &SequenceHandler{},
		StepHandler:     &StepHandler{},
		GeneralHandler:  NewGeneralHandler("1.0.0"),
	}
}

func TestNewRouter_HomePage(t *testing.T) {
	routes := setupRoutes()
	router := NewRouter(routes)

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")
	assert.Contains(t, rec.Header().Get("Content-Type"), "text/html", "Expected Content-Type text/html")
}

func TestNewRouter_OpenAPISpec_Success(t *testing.T) {
	routes := setupRoutes()
	router := NewRouter(routes)

	// Create a temporary OpenAPI file
	filePath := "internal/api/docs/openapi3.yaml"
	_ = os.MkdirAll("internal/api/docs", os.ModePerm)
	defer os.RemoveAll("internal")

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/docs/openapi3.yaml", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code, "Expected status code 200")
}

func TestNewRouter_OpenAPISpec_NotFound(t *testing.T) {
	routes := setupRoutes()
	router := NewRouter(routes)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/docs/openapi3.yaml", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code, "Expected status code 404 for missing OpenAPI spec")
}

func TestNewRouter_GeneralRoutes(t *testing.T) {
	routes := setupRoutes()
	router := NewRouter(routes)

	tests := []struct {
		path       string
		method     string
		statusCode int
	}{
		{"/api/v1/health", http.MethodGet, http.StatusOK},
		{"/api/v1/info", http.MethodGet, http.StatusOK},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, test.path, nil)
		if err != nil {
			t.Fatalf("Could not create request for %s: %v", test.path, err)
		}

		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, test.statusCode, rec.Code, "Expected status code %d for %s", test.statusCode, test.path)
	}
}
