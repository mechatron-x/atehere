package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type (
	CustomerSignUp struct {
		cs *service.Customer
	}

	CustomerProfile struct {
		cs *service.Customer
	}
)

func NewCustomerSignUp(customerService *service.Customer) CustomerSignUp {
	return CustomerSignUp{
		cs: customerService,
	}
}

func NewCustomerProfile(customerService *service.Customer) CustomerProfile {
	return CustomerProfile{
		cs: customerService,
	}
}

func (u CustomerSignUp) Pattern() string {
	return "POST /api/v1/customer/auth"
}

func (p CustomerProfile) Pattern() string {
	return "GET /api/v1/customer/profile"
}

func (u CustomerSignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.SignUpCustomer{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customer, err := u.cs.SignUp(reqBody.Customer)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.SignUpCustomer{Customer: customer}

	response.Encode(w, resp, http.StatusCreated)
}

func (p CustomerProfile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		errorHandler(w, err)
		return
	}

	customerProfile, err := p.cs.GetProfile(token)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.CustomerProfile{CustomerProfile: customerProfile}

	response.Encode(w, resp, http.StatusOK)
}
