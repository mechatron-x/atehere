package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type Customer struct {
	cs service.Customer
}

func NewCustomerHandler(cs service.Customer) Customer {
	return Customer{cs}
}

func (ch Customer) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.SignUpCustomer{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customer, err := ch.cs.SignUp(reqBody.Customer)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.SignUpCustomer{Customer: customer}

	response.Encode(w, resp, http.StatusCreated)
}

func (ch Customer) GetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		errorHandler(w, err)
		return
	}

	customerProfile, err := ch.cs.GetProfile(token)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.CustomerProfile{CustomerProfile: customerProfile}

	response.Encode(w, resp, http.StatusOK)
}

func (ch Customer) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		errorHandler(w, err)
		return
	}

	reqBody := &request.UpdateCustomer{}
	err = request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customerProfile, err := ch.cs.UpdateProfile(token, reqBody.Customer)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.CustomerProfile{CustomerProfile: customerProfile}

	response.Encode(w, resp, http.StatusOK)
}
