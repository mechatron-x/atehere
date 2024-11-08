package dto

type (
	TableCreate struct {
		Name string `json:"name"`
	}

	Table struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
