package dto

type (
	Session struct {
		SessionID string `json:"session_id"`
	}

	SessionState struct {
		AvailableStates []string `json:"available_session_states"`
		State           string   `json:"session_state"`
	}
)
