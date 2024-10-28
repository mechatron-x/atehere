package handler

import "net/http"

type Health struct{}

func NewHealth() Health {
	return Health{}
}

func (hh Health) GetHealth(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Healthy"))
}
