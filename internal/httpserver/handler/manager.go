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
	reqBody := &request.ManagerSignUp{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		response.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	manager, err := mh.ms.SignUp(reqBody.ManagerSignUp)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.ManagerSignUp{Manager: manager}
	response.Encode(w, resp, http.StatusCreated)
}

func (mh Manager) GetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	manager, err := mh.ms.GetProfile(token)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.ManagerGetProfile{Manager: manager}
	response.Encode(w, resp, http.StatusOK)
}

func (mh Manager) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	reqBody := &request.ManagerUpdateProfile{}
	err = request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	manager, err := mh.ms.UpdateProfile(token, reqBody.Manager)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.ManagerGetProfile{Manager: manager}
	response.Encode(w, resp, http.StatusOK)
}
