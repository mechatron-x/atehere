package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/header"
	"github.com/mechatron-x/atehere/internal/handler/request"
	"github.com/mechatron-x/atehere/internal/handler/response"
	"github.com/mechatron-x/atehere/internal/usermanagement/dto"
	"github.com/mechatron-x/atehere/internal/usermanagement/service"
)

type ManagerHandler struct {
	ms service.ManagerService
}

func NewManager(ms service.ManagerService) ManagerHandler {
	return ManagerHandler{ms}
}

func (mh ManagerHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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

func (mh ManagerHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
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

func (mh ManagerHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
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
