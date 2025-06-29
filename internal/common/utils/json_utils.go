package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func DecodeJSON[T any](r *http.Request, dst *T) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("content-type must be application/json")
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}


func EncodeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}