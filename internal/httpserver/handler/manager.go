package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type Manager struct {
	ms service.Manager
}

func NewManagerHandler(ms service.Manager) Manager {
	return Manager{ms}
}

func (mh Manager) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.SignUpManager{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	manager, err := mh.ms.SignUp(reqBody.Manager)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.SignUpManager{Manager: manager}

	response.Encode(w, resp, http.StatusCreated)
}

func (mh Manager) GetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		errorHandler(w, err)
		return
	}

	managerProfile, err := mh.ms.GetProfile(token)
	if err != nil {
		errorHandler(w, err)
		return
	}

	resp := response.ManagerProfile{ManagerProfile: managerProfile}

	response.Encode(w, resp, http.StatusOK)
}
