package handler

import "net/http"

type Health struct{}

func NewHealth() Router {
	return Health{}
}

// Pattern implements Router.
func (h Health) Pattern() string {
	return "GET /health"
}

// ServeHTTP implements Router.
func (h Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Healthy"))
}
