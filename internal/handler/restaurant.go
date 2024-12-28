package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/handler/header"
	"github.com/mechatron-x/atehere/internal/handler/request"
	"github.com/mechatron-x/atehere/internal/handler/response"
	"github.com/mechatron-x/atehere/internal/restaurant/dto"
	"github.com/mechatron-x/atehere/internal/restaurant/service"
)

type RestaurantHandler struct {
	rs service.RestaurantService
}

func NewRestaurant(rs service.RestaurantService) RestaurantHandler {
	return RestaurantHandler{rs}
}

func (rh RestaurantHandler) Create(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.RestaurantCreate{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		response.Encode(w, nil, err, http.StatusBadRequest)
		return
	}

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	restaurant, err := rh.rs.Create(token, reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, restaurant, nil, http.StatusCreated)
}

func (rh RestaurantHandler) GetOneForCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("restaurant_id")

	restaurantSummary, err := rh.rs.GetOneForCustomer(id)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	resp := &response.Restaurant[*dto.RestaurantSummary]{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurant:           restaurantSummary,
	}
	response.Encode(w, resp, nil)
}

func (rh RestaurantHandler) ListForManager(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	restaurants, err := rh.rs.ListForManager(token)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	resp := &response.RestaurantList[dto.Restaurant]{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurants:          restaurants,
	}
	response.Encode(w, resp, nil)
}

func (rh RestaurantHandler) ListForCustomer(w http.ResponseWriter, r *http.Request) {
	reqBody := &dto.RestaurantFilter{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	restaurants, err := rh.rs.ListForCustomer(reqBody)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	resp := &response.RestaurantList[dto.RestaurantSummary]{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurants:          restaurants,
	}
	response.Encode(w, resp, nil)
}

func (rh RestaurantHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("restaurant_id")

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	err = rh.rs.Delete(token, id)
	if err != nil {
		response.Encode(w, nil, err)
		return
	}

	response.Encode(w, nil, nil, http.StatusOK)
}
