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

func setupStepRouter(handler *StepHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/steps", handler.CreateStep).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/steps/{id}", handler.UpdateStep).Methods(http.MethodPut)
	router.HandleFunc("/api/v1/steps/{id}", handler.DeleteStep).Methods(http.MethodDelete)
	router.HandleFunc("/api/v1/steps", handler.ListSteps).Methods(http.MethodGet)
	return router
}

func TestCreateStep_Success(t *testing.T) {
	mockService := &StepServiceMock{
		CreateStepFunc: func(ctx context.Context, step *models.Step) (int64, error) {
			return 1, nil
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	step := models.Step{Subject: "Test Step"}
	body, _ := json.Marshal(step)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/steps", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Step created successfully", response["message"])
	assert.Equal(t, float64(1), response["data"].(map[string]interface{})["id"])
}

func TestCreateStep_InvalidBody(t *testing.T) {
	mockService := &StepServiceMock{}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/steps", bytes.NewReader([]byte("invalid json")))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid request body", response["message"])
}

func TestUpdateStep_Success(t *testing.T) {
	mockService := &StepServiceMock{
		UpdateStepFunc: func(ctx context.Context, step *models.Step) error {
			return nil
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	step := models.Step{Subject: "Updated Step"}
	body, _ := json.Marshal(step)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/steps/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Step updated successfully", response["message"])
}
func TestCreateStep_ServiceError(t *testing.T) {
	mockService := &StepServiceMock{
		CreateStepFunc: func(ctx context.Context, step *models.Step) (int64, error) {
			return 0, errors.New("service error")
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	step := models.Step{Subject: "Test Step"}
	body, _ := json.Marshal(step)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/steps", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Failed to create step", response["message"])
	assert.Equal(t, "service error", response["errors"])
}

func TestUpdateStep_InvalidID(t *testing.T) {
	mockService := &StepServiceMock{}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/steps/abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid query parameter", response["message"])
}

func TestDeleteStep_Success(t *testing.T) {
	mockService := &StepServiceMock{
		DeleteStepFunc: func(ctx context.Context, id int64) error {
			return nil
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/steps/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Step deleted successfully", response["message"])
}

func TestDeleteStep_InvalidID(t *testing.T) {
	mockService := &StepServiceMock{}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/steps/abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid query parameter", response["message"])
}

func TestListSteps_Success(t *testing.T) {
	mockService := &StepServiceMock{
		ListStepsFunc: func(ctx context.Context, sequenceID int64) ([]*models.Step, error) {
			return []*models.Step{{Subject: "Step 1"}, {Subject: "Step 2"}}, nil
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/steps?sequenceId=1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Steps fetched successfully", response["message"])
	assert.Len(t, response["data"].([]interface{}), 2)
}

func TestListSteps_InvalidSequenceID(t *testing.T) {
	mockService := &StepServiceMock{}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/steps?sequenceId=abc", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid query parameter", response["message"])
}

func TestUpdateStep_InvalidRequestBody(t *testing.T) {
	mockService := &StepServiceMock{}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/steps/1", bytes.NewReader([]byte("invalid json")))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Invalid request body", response["message"])
	assert.Contains(t, response["errors"], "invalid character")
}
func TestUpdateStep_ServiceError(t *testing.T) {
	mockService := &StepServiceMock{
		UpdateStepFunc: func(ctx context.Context, step *models.Step) error {
			return errors.New("service error")
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	step := models.Step{Subject: "Test Step"}
	body, _ := json.Marshal(step)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/steps/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Failed to update step", response["message"])
	assert.Equal(t, "service error", response["errors"])
}

func TestDeleteStep_ServiceError(t *testing.T) {
	mockService := &StepServiceMock{
		DeleteStepFunc: func(ctx context.Context, id int64) error {
			return errors.New("service error")
		},
	}
	handler := NewStepHandler(mockService)
	router := setupStepRouter(handler)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/steps/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	var response map[string]interface{}
	_ = json.NewDecoder(rec.Body).Decode(&response)
	assert.Equal(t, "Failed to delete step", response["message"])
	assert.Equal(t, "service error", response["errors"])
}
