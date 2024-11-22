package response

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
)

func Encode(w http.ResponseWriter, data any, err error, status ...int) {
	resp := Payload[any]{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(status) > 0 {
		w.WriteHeader(status[0])
	}

	if err != nil {
		resp.Error = newErrorData(err, status...)
	} else if data != nil {
		if reflect.TypeOf(data).Kind() == reflect.Pointer {
			resp.Data = data
		}
	} else {
		return
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func newErrorData(err error, httpStatus ...int) *ErrorData {
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

	return &ErrorData{
		Status:    status,
		Code:      code,
		Message:   message,
		CreatedAt: now,
	}
}
