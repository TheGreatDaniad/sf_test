package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sf_test/internal/core"
	"sf_test/internal/models"

	"github.com/gorilla/mux"
)

type SequenceHandler struct {
	sequenceService core.SequenceService
}

func NewSequenceHandler(service core.SequenceService) *SequenceHandler {
	return &SequenceHandler{sequenceService: service}
}

func (h *SequenceHandler) CreateSequence(w http.ResponseWriter, r *http.Request) {
	var sequence models.Sequence
	if err := json.NewDecoder(r.Body).Decode(&sequence); err != nil {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse(err.Error(), "Invalid request body"))
		return
	}

	id, err := h.sequenceService.CreateSequence(r.Context(), &sequence)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to create sequence"))
		return
	}

	WriteResponse(w, http.StatusCreated, SuccessResponse(map[string]int64{"id": id}, "Sequence created successfully"))
}

func (h *SequenceHandler) UpdateTracking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse("Invalid ID", "Invalid query parameter"))
		return
	}

	var payload struct {
		OpenTracking  bool `json:"openTracking"`
		ClickTracking bool `json:"clickTracking"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse(err.Error(), "Invalid request body"))
		return
	}
	err = h.sequenceService.UpdateTracking(r.Context(), id, payload.OpenTracking, payload.ClickTracking)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to update tracking"))
		return
	}

	WriteResponse(w, http.StatusOK, SuccessResponse(nil, "Tracking updated successfully"))
}

func (h *SequenceHandler) GetSequence(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		WriteResponse(w, http.StatusBadRequest, ErrorResponse("Invalid ID", "Invalid query parameter"))
		return
	}

	sequence, err := h.sequenceService.GetSequence(r.Context(), id)
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, ErrorResponse(err.Error(), "Failed to fetch sequence"))
		return
	}

	WriteResponse(w, http.StatusOK, SuccessResponse(sequence, "Sequence fetched successfully"))
}
