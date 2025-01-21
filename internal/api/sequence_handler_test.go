package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"sf_test/internal/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter(handler *SequenceHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/sequences", handler.CreateSequence).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/sequences/{id}", handler.UpdateTracking).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/sequences/{id}", handler.GetSequence).Methods(http.MethodGet)
	return router
}

func TestCreateSequence_Success(t *testing.T) {
	mockService := &SequenceServiceMock{
		CreateSequenceFunc: func(ctx context.Context, sequence *models.Sequence) (int64, error) {
			return 1, nil
		},
	}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	sequence := models.Sequence{Name: "Test Sequence"}
	body, _ := json.Marshal(sequence)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/sequences", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Sequence created successfully", response["message"])
	assert.Equal(t, float64(1), response["data"].(map[string]interface{})["id"])
}

func TestCreateSequence_InvalidBody(t *testing.T) {
	mockService := &SequenceServiceMock{}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/sequences", bytes.NewReader([]byte("invalid json")))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid request body", response["message"])
}

func TestUpdateTracking_Success(t *testing.T) {
	mockService := &SequenceServiceMock{
		UpdateTrackingFunc: func(ctx context.Context, id int64, openTracking, clickTracking bool) error {
			return nil
		},
	}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	payload := map[string]bool{
		"openTracking":  true,
		"clickTracking": false,
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/sequences/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Tracking updated successfully", response["message"])
}

func TestUpdateTracking_InvalidID(t *testing.T) {
	mockService := &SequenceServiceMock{}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/sequences/abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid query parameter", response["message"])
}

func TestGetSequence_Success(t *testing.T) {
	mockService := &SequenceServiceMock{
		GetSequenceFunc: func(ctx context.Context, id int64) (*models.Sequence, error) {
			return &models.Sequence{Name: "Test Sequence"}, nil
		},
	}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sequences/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Sequence fetched successfully", response["message"])
}

func TestGetSequence_InvalidID(t *testing.T) {
	mockService := &SequenceServiceMock{}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sequences/abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid query parameter", response["message"])
}

func TestGetSequence_NotFound(t *testing.T) {
	mockService := &SequenceServiceMock{
		GetSequenceFunc: func(ctx context.Context, id int64) (*models.Sequence, error) {
			return nil, errors.New("not found")
		},
	}
	handler := NewSequenceHandler(mockService)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sequences/99", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Failed to fetch sequence", response["message"])
}
