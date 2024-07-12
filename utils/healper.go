package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseJson(w http.ResponseWriter, r *http.Request, v any) error {
	if r.Body == nil {

		return errors.New("invalid payload")
	}
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return errors.New("invalid payload")
	}
	return nil
}

// func WriteError(w http.ResponseWriter, err error, code int)  {
// 	w.WriteHeader(code)
//   map[string]error{"error": err}

// }

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
