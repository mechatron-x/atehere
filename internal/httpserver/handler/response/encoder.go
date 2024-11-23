package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
)

func Encode(w http.ResponseWriter, data any, err error, status ...int) {
	if err != nil {
		encodeError(w, err, status...)
		return
	}
	if data != nil {
		encodeData(w, data, status...)
		return
	}
}

func encodeError(w http.ResponseWriter, err error, status ...int) {
	defaultStatus := http.StatusInternalServerError
	if len(status) > 0 {
		defaultStatus = status[0]
	}

	code := "unhandled_code"
	message := err.Error()
	if domainErr, ok := err.(core.DomainError); ok {
		code = domainErr.Code()
		defaultStatus = mapDomainErrorToStatus(code, defaultStatus)
	}

	payload := &Payload[any]{
		Error: &Error{
			Status:    defaultStatus,
			Code:      code,
			Message:   message,
			CreatedAt: time.Now(),
		},
	}

	w.Header().Set(header.ContentTypeKey, header.ContentTypeJSON)
	w.WriteHeader(defaultStatus)
	_ = json.NewEncoder(w).Encode(payload)
}

func encodeData(w http.ResponseWriter, data any, status ...int) {
	defaultStatus := http.StatusOK
	if len(status) > 0 {
		defaultStatus = status[0]
	}

	payload := &Payload[any]{
		Data: data,
	}

	w.Header().Set(header.ContentTypeKey, header.ContentTypeJSON)
	w.WriteHeader(defaultStatus)
	_ = json.NewEncoder(w).Encode(payload)
}

func mapDomainErrorToStatus(code string, defaultStatus int) int {
	switch code {
	case core.CodeUnauthorized:
		return http.StatusUnauthorized
	case core.CodePersistenceFailure:
		return http.StatusInternalServerError
	case core.CodeResourceNotFound:
		return http.StatusNotFound
	case core.CodeDataConflict:
		return http.StatusConflict
	case core.CodeValidationFailure, core.CodeDomainIntegrityViolation:
		return http.StatusBadRequest
	default:
		return defaultStatus
	}
}
