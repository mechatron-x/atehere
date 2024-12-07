package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type CustomerHandler struct {
	cs service.CustomerService
}

func NewCustomer(cs service.CustomerService) CustomerHandler {
	return CustomerHandler{cs}
}

func (ch CustomerHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.CustomerSignUp{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		response.Encode(w, nil, err, http.StatusBadRequest)
		return
	}

	customer, err := ch.cs.SignUp(reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, customer, nil, http.StatusCreated)
}

func (ch CustomerHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	customer, err := ch.cs.GetProfile(token)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, customer, err)
}

func (ch CustomerHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	reqBody := &dto.Customer{}
	err = request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customer, err := ch.cs.UpdateProfile(token, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, customer, nil)
}
