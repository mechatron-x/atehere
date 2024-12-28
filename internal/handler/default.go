package handler

import (
	"errors"
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/response"
)

type DefaultHandler struct {
}

func NewDefault() DefaultHandler {
	return DefaultHandler{}
}

func (dh DefaultHandler) NoHandler(w http.ResponseWriter, r *http.Request) {
	response.Encode(w, nil, errors.New("no handler found"), http.StatusForbidden)
}
