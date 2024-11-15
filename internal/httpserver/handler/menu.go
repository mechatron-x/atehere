package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/menu/service"
)

type Menu struct {
	ms service.Menu
}

func NewMenuHandler(ms service.Menu) Menu {
	return Menu{ms: ms}
}

func (mh Menu) Create(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.MenuCreate{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	menu, err := mh.ms.Create(token, reqBody.MenuCreate)
	if err != nil {
		response.EncodeError(w, err)
	}

	resp := response.MenuCreate{Menu: menu}
	response.Encode(w, resp, http.StatusCreated)
}

func (mh Menu) GetMenu(w http.ResponseWriter, r *http.Request) {
	restaurantID := r.PathValue("id")
	category := r.PathValue("category")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	menu, err := mh.ms.GetMenuByCategory(token, restaurantID, category)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.Menu{
		Menu: *menu,
	}
	response.Encode(w, resp, http.StatusOK)
}
