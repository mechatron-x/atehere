package response

type (
	MenuList[TDTO any] struct {
		Menus []TDTO `json:"menus"`
	}
)
