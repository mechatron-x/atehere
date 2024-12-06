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

func (sh Session) PlaceOrders(w http.ResponseWriter, r *http.Request) {
	tableID := r.PathValue("table_id")

	reqBody := &dto.PlaceOrders{}
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

	err = sh.ss.PlaceOrders(token, tableID, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}
}

func (sh Session) Checkout(w http.ResponseWriter, r *http.Request) {
	tableID := r.PathValue("table_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	err = sh.ss.Checkout(token, tableID)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}
}

func (sh Session) CustomerOrdersView(w http.ResponseWriter, r *http.Request) {
	tableID := r.PathValue("table_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	orders, err := sh.ss.CustomerOrdersView(token, tableID)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, orders, nil)
}

func (sh Session) ManagerOrdersView(w http.ResponseWriter, r *http.Request) {
	table_id := r.PathValue("table_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	orders, err := sh.ss.ManagerOrdersView(token, table_id)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, orders, nil)
}

func (sh Session) TableOrdersView(w http.ResponseWriter, r *http.Request) {
	table_id := r.PathValue("table_id")

	orders, err := sh.ss.TableOrdersView(table_id)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, orders, nil)
}
