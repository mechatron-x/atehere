package response

type (
	Health struct {
		Uptime string `json:"uptime,omitempty"`
		Status string `json:"status,omitempty"`
	}
)
