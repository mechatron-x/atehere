package handler

import (
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/header"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/request"
	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
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
	}

	resp := response.RestaurantCreate{Restaurant: restaurant}
	response.Encode(w, resp, http.StatusCreated)
}

func (rh Restaurant) List(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	restaurants, err := rh.rs.List(page)
	if err != nil {
		response.EncodeError(w, err)
		return
	}

	resp := response.RestaurantList{
		AvailableWorkingDays: rh.rs.AvailableWorkingDays(),
		FoundationYearFormat: rh.rs.FoundationYearFormat(),
		WorkingTimeFormat:    rh.rs.WorkingTimeFormat(),
		Restaurants:          restaurants,
	}
	response.Encode(w, resp, http.StatusOK)
}
