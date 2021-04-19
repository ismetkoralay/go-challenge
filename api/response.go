package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %v", e.StatusCode, e.Message)
}

type SuccessResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type ServerError interface {
	GetStatusCode() int
	GetMessage() string
}

func JsonResponse(w http.ResponseWriter, statusCode int, isSuccess bool, response interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	if isSuccess {
		encoder.Encode(response)
	} else {
		if message == "" {
			message = http.StatusText(statusCode)
		}
		encoder.Encode(ErrorResponse{
			statusCode,
			message,
		})
	}
}
