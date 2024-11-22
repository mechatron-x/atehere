package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
)

func Encode(w http.ResponseWriter, data any, err error, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	responseStatus := determineStatus(err, status)

	var payload Payload[any]

	if err != nil {
		payload.Error = createError(err, responseStatus)
		responseStatus = payload.Error.Status
	} else if data != nil {
		payload.Data = data
	}

	w.WriteHeader(responseStatus)
	_ = json.NewEncoder(w).Encode(payload)
}

func determineStatus(err error, status []int) int {
	if err != nil {
		return http.StatusInternalServerError
	}
	if len(status) > 0 {
		return status[0]
	}
	return http.StatusOK
}

func createError(err error, defaultStatus int) *Error {
	status := defaultStatus
	code := "unhandled_error"
	message := err.Error()

	if domainErr, ok := err.(core.DomainError); ok {
		code = domainErr.Code()
		status = mapDomainErrorToStatus(domainErr.Code(), status)
	}

	return &Error{
		Status:    status,
		Code:      code,
		Message:   message,
		CreatedAt: time.Now(),
	}
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
	case core.CodeValidationFailure:
		return http.StatusBadRequest
	default:
		return defaultStatus
	}
}
