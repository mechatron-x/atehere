package handler

import (
	"errors"
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

type Default struct {
}

func NewDefault() Default {
	return Default{}
}

func (dh Default) NoHandler(w http.ResponseWriter, r *http.Request) {
	response.Encode(w, nil, errors.New("no handler found"), http.StatusForbidden)
}
