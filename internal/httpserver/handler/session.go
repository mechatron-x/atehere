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
	tableID := r.PathValue("table_id")

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

	err = sh.ss.PlaceOrder(token, tableID, reqBody)
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

func (sh Session) CustomerOrders(w http.ResponseWriter, r *http.Request) {
	tableID := r.PathValue("table_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	orders, err := sh.ss.CustomerOrders(token, tableID)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	resp := &response.OrderList[dto.OrderCustomerView]{
		Orders: orders,
	}
	response.Encode(w, resp, nil)
}

func (sh Session) TableOrders(w http.ResponseWriter, r *http.Request) {
	table_id := r.PathValue("table_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	orders, err := sh.ss.TableOrders(token, table_id)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	resp := &response.OrderList[dto.OrderTableView]{
		Orders: orders,
	}
	response.Encode(w, resp, nil)
}
