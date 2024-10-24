package response

type Error struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}
