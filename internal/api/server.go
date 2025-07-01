package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Handler[I any, O any] interface {
	Handle(ctx context.Context, req *I) (*O, error)
}

func MakeHTTPHandler[I any, O any](h Handler[I, O]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req I
		if err := DecodeJSON(r, &req); err != nil {
			EncodeJSON(w, http.StatusBadRequest, "invalid request body")
			return
		}

		resp, err := h.Handle(r.Context(), &req)
		if err != nil {
			EncodeJSON(w, http.StatusInternalServerError, "unexpected error")
			return
		}

		EncodeJSON(w, http.StatusOK, resp)
	}
}

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