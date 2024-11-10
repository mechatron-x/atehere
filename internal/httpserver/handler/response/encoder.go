package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
)

type (
	ErrorResponse struct {
		Status    int    `json:"status"`
		Code      string `json:"code"`
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

func EncodeError(w http.ResponseWriter, err error, httpStatus ...int) {
	status := http.StatusInternalServerError
	code := "unhandled"
	message := err.Error()
	now := time.Now().String()

	if len(httpStatus) > 0 {
		status = httpStatus[0]
	}

	if domainErr, ok := err.(core.DomainError); ok {
		code = domainErr.Code()

		switch domainErr.Code() {
		case core.CodeUnauthorized:
			status = http.StatusUnauthorized
		case core.CodePersistenceFailure:
			status = http.StatusInternalServerError
		case core.CodeResourceNotFound:
			status = http.StatusNotFound
		case core.CodeDataConflict:
			status = http.StatusConflict
		case core.CodeValidationFailure:
			status = http.StatusBadRequest
		}
	}

	errResp := ErrorResponse{
		Status:    status,
		Code:      code,
		Message:   message,
		CreatedAt: now,
	}

	Encode(w, errResp, status)
}
