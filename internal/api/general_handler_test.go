package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestNewGeneralHandler(t *testing.T) {
	version := "1.0.0"
	handler := NewGeneralHandler(version)

	if handler.version != version {
		t.Errorf("Expected version %s, got %s", version, handler.version)
	}

	if time.Since(handler.startTime) > time.Second {
		t.Errorf("Handler startTime is not initialized correctly")
	}
}

func TestHealthCheck(t *testing.T) {
	handler := NewGeneralHandler("1.0.0")
	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler.HealthCheck(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

}

func TestGetAPIInfo(t *testing.T) {
	handler := NewGeneralHandler("1.0.0")
	req, err := http.NewRequest("GET", "/api/v1/info", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler.GetAPIInfo(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}



	info, ok := body["data"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected data to be a map")
	}

	if info["version"] != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got %v", info["version"])
	}
}

func TestHomePage(t *testing.T) {
	handler := NewGeneralHandler("1.0.0")
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler.HomePage(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("Expected content type text/html, got %v", contentType)
	}
}

func TestGetEnvironment(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	defer os.Unsetenv("APP_ENV")

	env := getEnvironment()
	if env != "production" {
		t.Errorf("Expected environment 'production', got %v", env)
	}

	os.Unsetenv("APP_ENV")
	env = getEnvironment()
	if env != "development" {
		t.Errorf("Expected environment 'development', got %v", env)
	}
}

