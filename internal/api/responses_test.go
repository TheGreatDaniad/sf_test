package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteResponse_Success(t *testing.T) {
	// Arrange
	rec := httptest.NewRecorder()
	data := map[string]string{"key": "value"}
	message := "Request successful"
	response := SuccessResponse(data, message)

	// Act
	WriteResponse(rec, http.StatusOK, response)

	// Assert
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got %s", contentType)
	}

	var body APIResponse
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if body.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, body.Message)
	}

	if body.Data.(map[string]interface{})["key"] != "value" {
		t.Errorf("Expected data 'key: value', got %v", body.Data)
	}
}

func TestWriteResponse_Error(t *testing.T) {
	// Arrange
	rec := httptest.NewRecorder()
	errors := map[string]string{"error": "Invalid input"}
	message := "Request failed"
	response := ErrorResponse(errors, message)

	// Act
	WriteResponse(rec, http.StatusBadRequest, response)

	// Assert
	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got %s", contentType)
	}

	var body APIResponse
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if body.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, body.Message)
	}

	if body.Errors.(map[string]interface{})["error"] != "Invalid input" {
		t.Errorf("Expected error 'error: Invalid input', got %v", body.Errors)
	}
}

func TestSuccessResponse(t *testing.T) {
	// Arrange
	data := map[string]string{"key": "value"}
	message := "Success message"

	// Act
	response := SuccessResponse(data, message)

	// Assert
	if response.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, response.Message)
	}

	if response.Data.(map[string]string)["key"] != "value" {
		t.Errorf("Expected data 'key: value', got %v", response.Data)
	}

	if response.Errors != nil {
		t.Errorf("Expected no errors, got %v", response.Errors)
	}
}

func TestErrorResponse(t *testing.T) {
	// Arrange
	errors := map[string]string{"error": "Something went wrong"}
	message := "Error message"

	// Act
	response := ErrorResponse(errors, message)

	// Assert
	if response.Message != message {
		t.Errorf("Expected message '%s', got '%s'", message, response.Message)
	}

	if response.Errors.(map[string]string)["error"] != "Something went wrong" {
		t.Errorf("Expected error 'error: Something went wrong', got %v", response.Errors)
	}

	if response.Data != nil {
		t.Errorf("Expected no data, got %v", response.Data)
	}
}
