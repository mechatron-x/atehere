package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type (
	CustomerSignUp struct {
		cs service.Customer
	}
)

func NewUserSignUp(customerService service.Customer) CustomerSignUp {
	return CustomerSignUp{
		cs: customerService,
	}
}

func (u CustomerSignUp) Pattern() string {
	return "POST /api/v1/customer/auth"
}

func (u CustomerSignUp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.SignUpCustomer{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	customer, err := u.cs.SignUp(reqBody.CustomerSignUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := response.SignUpCustomer{CustomerSignUp: customer}

	response.Encode(w, userResponse, http.StatusCreated)
}
