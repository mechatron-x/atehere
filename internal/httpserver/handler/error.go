package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

func errorHandler(w http.ResponseWriter, err error) {
	now := time.Now().String()
	domainErr, ok := err.(core.DomainError)
	if !ok {
		responseErr := &response.Error{
			Code:      http.StatusInternalServerError,
			Message:   "internal server error",
			CreatedAt: now,
		}
		response.Encode(w, responseErr, http.StatusInternalServerError)
		return
	}

	var code int

	var validationErr *core.ValidationError
	var persistenceErr *core.PersistenceError
	var conflictErr *core.ConflictError

	if errors.As(domainErr, &validationErr) {
		code = http.StatusBadRequest
	} else if errors.As(domainErr, &persistenceErr) {
		code = http.StatusInternalServerError
	} else if errors.As(domainErr, &conflictErr) {
		code = http.StatusConflict
	}

	responseErr := &response.Error{
		Code:      code,
		Message:   domainErr.Reason().Error(),
		CreatedAt: now,
	}

	response.Encode(w, responseErr, code)
}
