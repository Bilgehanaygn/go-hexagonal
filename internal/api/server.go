package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler[I any, O any] interface {
	Handle(ctx context.Context, req *I) (*O, error)
}

func MakeHTTPHandler[I any, O any](h Handler[I, O]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1) create an empty request struct
		var req I

		if err := bindURLParams(r, &req); err != nil {
			EncodeJSON(w, http.StatusBadRequest, "invalid path parameter: "+err.Error())
			return
		}

		if r.Header.Get("Content-Type") == "application/json" {
			if err := DecodeJSON(r, &req); err != nil {
				EncodeJSON(w, http.StatusBadRequest, "invalid request body")
				return
			}
		}

		resp, err := h.Handle(r.Context(), &req)
		if err != nil {
			EncodeJSON(w, http.StatusInternalServerError, "unexpected error")
			return
		}

		EncodeJSON(w, http.StatusOK, resp)
	}
}

func bindURLParams[T any](r *http.Request, dst *T) error {
	// dest must be a pointer to a struct
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("destination must be pointer to struct")
	}
	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		tag := sf.Tag.Get("param")
		if tag == "" {
			continue
		}

		parts := strings.Split(tag, ",")
		paramName := parts[0]
		required := len(parts) > 1 && parts[1] == "required"

		raw := chi.URLParam(r, paramName)
		if raw == "" {
			if required {
				return fmt.Errorf("required path parameter '%s' is missing", paramName)
			}
			continue
		}

		fv := v.Field(i)
		switch sf.Type.Kind() {
		case reflect.String:
			fv.SetString(raw)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(raw, 10, 64)
			if err != nil {
				return err
			}
			fv.SetInt(n)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			u, err := strconv.ParseUint(raw, 10, 64)
			if err != nil {
				return err
			}
			fv.SetUint(u)
		default:
			// special‐case known types e.g. uuid.UUID
			if sf.Type == reflect.TypeOf(uuid.UUID{}) {
				id, err := uuid.Parse(raw)
				if err != nil {
					return err
				}
				fv.Set(reflect.ValueOf(id))
			}
			// add more conversions here as needed
		}
	}

	return nil
}

// rest of your helpers unchanged:

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
