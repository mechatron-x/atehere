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
	reqBody := &request.CustomerSignUp{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customer, err := ch.cs.SignUp(reqBody.CustomerSignUp)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.CustomerSignUp{Customer: customer}
	response.Encode(w, resp, http.StatusCreated)
}

func (ch Customer) GetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	customer, err := ch.cs.GetProfile(token)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.CustomerGetProfile{Customer: customer}
	response.Encode(w, resp, http.StatusOK)
}

func (ch Customer) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	reqBody := &request.CustomerUpdateProfile{}
	err = request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customer, err := ch.cs.UpdateProfile(token, reqBody.Customer)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.CustomerUpdateProfile{Customer: customer}
	response.Encode(w, resp, http.StatusOK)
}
