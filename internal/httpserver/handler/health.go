package handler

import "net/http"

type Health struct{}

func NewHealth() Route {
	return &Health{}
}

// Pattern implements Route.
func (h *Health) Pattern() string {
	return "GET /health"
}

// ServeHTTP implements Route.
func (h *Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}
