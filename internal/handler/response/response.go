package response

import "time"

type (
	Error struct {
		Status    int       `json:"status"`
		Code      string    `json:"code"`
		Message   string    `json:"message"`
		CreatedAt time.Time `json:"created_at"`
	}

	Payload[TData any] struct {
		Data  TData  `json:"data,omitempty"`
		Error *Error `json:"error,omitempty"`
	}
)
