package utils

import (
	"net/http"
	"errors"
	"encoding/json"
)

func ParseJSON(r http.Request, payload any) error {

	if r.Body == nil {
		return errors.New("request body is empty")
	}
	return json.NewDecoder(r.Body).Decode(&payload)	
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	
	type errorResponse struct {
		Error string `json:"error"`
	}

	WriteJSON(w, status, errorResponse{Error: err.Error()})
}