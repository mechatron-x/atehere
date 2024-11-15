package response

type (
	ErrorData struct {
		Status    int    `json:"status"`
		Code      string `json:"code"`
		Message   string `json:"message"`
		CreatedAt string `json:"created_at"`
	}

	Payload[TData any] struct {
		Data  TData      `json:"data,omitempty"`
		Error *ErrorData `json:"error,omitempty"`
	}
)
