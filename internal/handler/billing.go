package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/billing/dto"
	"github.com/mechatron-x/atehere/internal/billing/service"
	"github.com/mechatron-x/atehere/internal/handler/header"
	"github.com/mechatron-x/atehere/internal/handler/request"
	"github.com/mechatron-x/atehere/internal/handler/response"
)

type BillingHandler struct {
	service *service.BillingService
}

func NewBilling(service *service.BillingService) BillingHandler {
	return BillingHandler{
		service: service,
	}
}

func (rcv BillingHandler) GetBill(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	sessionID := r.PathValue("session_id")

	bill, err := rcv.service.Get(token, sessionID)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, bill, nil)
}

func (rcv BillingHandler) Pay(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.PayBillItems{}
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

	sessionID := r.PathValue("session_id")

	err = rcv.service.Pay(token, sessionID, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}
}

func (rcv BillingHandler) GetPastBills(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	pastBills, err := rcv.service.PastBills(token)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, pastBills, nil)
}
