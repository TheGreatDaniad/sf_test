package api

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
}

// WriteResponse sends a JSON response with the given status code.
func WriteResponse(w http.ResponseWriter, status int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// SuccessResponse generates a successful response.
func SuccessResponse(data interface{}, message string) APIResponse {
	return APIResponse{
		Data:    data,
		Message: message,
	}
}

// ErrorResponse generates an error response.
func ErrorResponse(errors interface{}, message string) APIResponse {
	return APIResponse{
		Errors:  errors,
		Message: message,
	}
}
