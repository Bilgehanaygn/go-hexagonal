package api

import (
	"context"
	"net/http"
	"strings"
)


type Handler[I any, O any] interface {
	Handle(context.Context, *I) (*O, error)
}

func Serve[I any, O any](path string, h Handler[I, O]) {
    parts := strings.SplitN(path, " ", 2)

	if len(parts) != 2 {
        panic("api.Serve: route must be in form METHOD /path")
	}
	method, pattern := parts[0], parts[1]

    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request)) {
        if r.Method != method {
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
            return
        }

		var req I
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

		res, err := h.Handle(r.Context(), &req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

		w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(res); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
	}
}

func (s *Server) Start(){}
func (s* Server) Stop(){}