package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

func errorHandler(w http.ResponseWriter, err error) {
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

	responseErr := &response.Error{
		Code:      code,
		Message:   err.Error(),
		CreatedAt: now,
	}

	response.Encode(w, responseErr, code)
}
