package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/session/dto"
	"github.com/mechatron-x/atehere/internal/session/service"
)

type SessionHandler struct {
	ss service.SessionService
}

func NewSession(ss service.SessionService) SessionHandler {
	return SessionHandler{ss}
}

func (sh SessionHandler) PlaceOrders(w http.ResponseWriter, r *http.Request) {
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

	session, err := sh.ss.PlaceOrders(token, tableID, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, session, nil)
}

func (sh SessionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	tableID := r.PathValue("table_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	session, err := sh.ss.Checkout(token, tableID)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, session, nil, http.StatusAccepted)
}

func (sh SessionHandler) CustomerOrdersView(w http.ResponseWriter, r *http.Request) {
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

func (sh SessionHandler) ManagerOrdersView(w http.ResponseWriter, r *http.Request) {
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

func (sh SessionHandler) TableOrdersView(w http.ResponseWriter, r *http.Request) {
	table_id := r.PathValue("table_id")

	orders, err := sh.ss.TableOrdersView(table_id)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, orders, nil)
}
