package response

type (
	OrderList[TDTO any] struct {
		Orders []TDTO `json:"orders"`
	}
)
