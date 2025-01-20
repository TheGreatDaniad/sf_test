package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sf_test/internal/core"
	"sf_test/internal/models"

	"github.com/gorilla/mux"
)

type StepHandler struct {
	stepService core.StepService
}

func NewStepHandler(service core.StepService) *StepHandler {
	return &StepHandler{stepService: service}
}

func (h *StepHandler) CreateStep(w http.ResponseWriter, r *http.Request) {
	var step models.Step
	if err := json.NewDecoder(r.Body).Decode(&step); err != nil {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse(err.Error(), "Invalid request body"))
		return
	}

	id, err := h.stepService.CreateStep(r.Context(), &step)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to create step"))
		return
	}

	WriteResponse(w, http.StatusCreated, SuccessResponse(map[string]int64{"id": id}, "Step created successfully"))
}

func (h *StepHandler) UpdateStep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse("Invalid ID", "Invalid query parameter"))
		return
	}
	var step models.Step
	if err := json.NewDecoder(r.Body).Decode(&step); err != nil {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse(err.Error(), "Invalid request body"))
		return
	}
	step.ID = id

	err = h.stepService.UpdateStep(r.Context(), &step)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to update step"))
		return
	}

	WriteResponse(w, http.StatusOK, SuccessResponse(nil, "Step updated successfully"))
}

func (h *StepHandler) DeleteStep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse("Invalid ID", "Invalid query parameter"))
		return
	}

	err = h.stepService.DeleteStep(r.Context(), id)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to delete step"))
		return
	}

	WriteResponse(w, http.StatusOK, SuccessResponse(nil, "Step deleted successfully"))
}

func (h *StepHandler) ListSteps(w http.ResponseWriter, r *http.Request) {
	sequenceIDStr := r.URL.Query().Get("sequenceId")
	sequenceID, err := strconv.ParseInt(sequenceIDStr, 10, 64)
	if err != nil || sequenceID <= 0 {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse("Invalid sequenceId", "Invalid query parameter"))
		return
	}

	steps, err := h.stepService.ListSteps(r.Context(), sequenceID)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to fetch steps"))
		return
	}

	WriteResponse(w, http.StatusOK, SuccessResponse(steps, "Steps fetched successfully"))
}
