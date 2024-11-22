package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/session/service"
)

type Session struct {
	ss service.Session
}

func NewSessionHandler(ss service.Session) Session {
	return Session{ss}
}

func (sh Session) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	table_id := r.PathValue("table_id")

	reqBody := &dto.OrderCreate{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		response.Encode(w, nil, err, http.StatusBadRequest)
		return
	}

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	err = sh.ss.PlaceOrder(token, table_id, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}
}
