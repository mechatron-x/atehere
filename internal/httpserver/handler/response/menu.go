package response

import "github.com/mechatron-x/atehere/internal/menu/dto"

type (
	MenuCreate struct {
		*dto.Menu
	}

	Menu struct {
		Menu dto.Menu `json:"menu"`
	}
)
