package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type Customer struct {
	cs service.Customer
}

func NewCustomerHandler(cs service.Customer) Customer {
	return Customer{cs}
}

func (ch Customer) SignUp(w http.ResponseWriter, r *http.Request) {
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

func (ch Customer) GetProfile(w http.ResponseWriter, r *http.Request) {
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

func (ch Customer) UpdateProfile(w http.ResponseWriter, r *http.Request) {
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
