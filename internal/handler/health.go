package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/response"
)

type HealthHandler struct{}

func NewHealth() HealthHandler {
	return HealthHandler{}
}

func (hh HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	resp := &response.Health{
		Status: "Healthy",
	}

	response.Encode(w, resp, nil)
}
