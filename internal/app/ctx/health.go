package ctx

import "github.com/mechatron-x/atehere/internal/httpserver/handler"

type Health struct {
	handler handler.Health
}

func NewHealth() Health {
	handler := handler.NewHealth()

	return Health{
		handler: handler,
	}
}

func (h Health) Handler() handler.Health {
	return h.handler
}
