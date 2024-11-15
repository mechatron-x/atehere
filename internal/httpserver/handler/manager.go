package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type Manager struct {
	ms service.Manager
}

func NewManagerHandler(ms service.Manager) Manager {
	return Manager{ms}
}

func (mh Manager) SignUp(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.ManagerSignUp{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		response.Encode(w, nil, err, http.StatusBadRequest)
		return
	}

	manager, err := mh.ms.SignUp(reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, manager, nil, http.StatusCreated)
}

func (mh Manager) GetProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	manager, err := mh.ms.GetProfile(token)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, manager, nil)
}

func (mh Manager) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	reqBody := &dto.Manager{}
	err = request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	manager, err := mh.ms.UpdateProfile(token, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, manager, nil)
}
