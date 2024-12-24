package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/header"
	"github.com/mechatron-x/atehere/internal/handler/request"
	"github.com/mechatron-x/atehere/internal/handler/response"
	"github.com/mechatron-x/atehere/internal/menu/dto"
	"github.com/mechatron-x/atehere/internal/menu/service"
)

type MenuHandler struct {
	ms service.MenuService
}

func NewMenu(ms service.MenuService) MenuHandler {
	return MenuHandler{ms: ms}
}

func (mh MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.MenuCreate{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err, http.StatusUnauthorized)
		return
	}

	menu, err := mh.ms.Create(token, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, menu, nil, http.StatusCreated)
}

func (mh MenuHandler) AddMenuItem(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.MenuItemCreate{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	reqBody.MenuID = r.PathValue("menu_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err, http.StatusUnauthorized)
		return
	}

	menu, err := mh.ms.AddMenuItem(token, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, menu, nil)
}

func (mh MenuHandler) ListForCustomer(w http.ResponseWriter, r *http.Request) {
	restaurantID := r.PathValue("restaurant_id")
	menuFilter := &dto.MenuFilter{
		RestaurantID: restaurantID,
	}

	menus, err := mh.ms.ListForCustomer(menuFilter)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	resp := &response.MenuList[dto.Menu]{
		Menus: menus,
	}
	response.Encode(w, resp, nil)
}
