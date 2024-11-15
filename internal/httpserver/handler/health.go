package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

type Health struct{}

func NewHealth() Health {
	return Health{}
}

func (hh Health) GetHealth(w http.ResponseWriter, r *http.Request) {
	resp := &response.Health{
		Status: "Healthy",
	}

	response.Encode(w, resp, nil)
}
