package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
	"github.com/mechatron-x/atehere/internal/restaurant/dto"
	"github.com/mechatron-x/atehere/internal/restaurant/service"
)

type Restaurant struct {
	rs service.Restaurant
}

func NewRestaurantHandler(rs service.Restaurant) Restaurant {
	return Restaurant{rs}
}

func (rh Restaurant) Create(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.RestaurantCreate{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		response.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	restaurant, err := rh.rs.Create(token, reqBody.RestaurantCreate)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.RestaurantCreate{Restaurant: restaurant}
	response.Encode(w, resp, http.StatusCreated)
}

func (rh Restaurant) GetOneForCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	restaurantSummary, err := rh.rs.GetOneForCustomer(id)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.Restaurant[*dto.RestaurantSummary]{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurant:           restaurantSummary,
	}
	response.Encode(w, resp, http.StatusOK)
}

func (rh Restaurant) ListForManager(w http.ResponseWriter, r *http.Request) {
	token, err := header.GetBearerToken(r.Header)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	restaurants, err := rh.rs.ListForManager(token)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.RestaurantList[dto.Restaurant]{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurants:          restaurants,
	}
	response.Encode(w, resp, http.StatusOK)
}

func (rh Restaurant) ListForCustomer(w http.ResponseWriter, r *http.Request) {
	reqBody := &request.RestaurantFilter{}
	err := request.Decode(r, w, reqBody)
	if err != nil {
		return
	}

	restaurants, err := rh.rs.ListForCustomer(reqBody.Page, reqBody.RestaurantFilter)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.RestaurantList[dto.RestaurantSummary]{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurants:          restaurants,
	}
	response.Encode(w, resp, http.StatusOK)
}
