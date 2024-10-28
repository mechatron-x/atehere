package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
)

type (
	ErrorResponse struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		CreatedAt string `json:"created_at"`
	}
)

func Encode(w http.ResponseWriter, v any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

func EncodeError(w http.ResponseWriter, err error) {
	now := time.Now().String()
	code := http.StatusInternalServerError

	if errors.Is(err, core.ErrConflict) {
		code = http.StatusConflict
	} else if errors.Is(err, core.ErrNotFound) {
		code = http.StatusNotFound
	} else if errors.Is(err, core.ErrPersistence) {
		code = http.StatusInternalServerError
	} else if errors.Is(err, core.ErrUnauthorized) {
		code = http.StatusUnauthorized
	} else if errors.Is(err, core.ErrValidation) {
		code = http.StatusBadRequest
	} else if errors.Is(err, header.ErrInvalidBearerToken) {
		code = http.StatusUnauthorized
	}

	errResp := ErrorResponse{
		Code:      code,
		Message:   err.Error(),
		CreatedAt: now,
	}

	Encode(w, errResp, code)
}
