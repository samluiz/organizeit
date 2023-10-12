package common

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, err *Error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
}

func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) ([]byte, error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return response, nil
}