package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/billing/dto"
	"github.com/mechatron-x/atehere/internal/billing/service"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

type BillingHandler struct {
	service *service.BillingService
}

func NewBilling(service *service.BillingService) BillingHandler {
	return BillingHandler{
		service: service,
	}
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
