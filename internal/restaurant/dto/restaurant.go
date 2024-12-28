package dto

import (
	"slices"
	"strings"
	"time"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/entity"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
)

type (
	RestaurantCreate struct {
		Name           string        `json:"name"`
		FoundationYear string        `json:"foundation_year"`
		PhoneNumber    string        `json:"phone_number"`
		OpeningTime    string        `json:"opening_time"`
		ClosingTime    string        `json:"closing_time"`
		Image          string        `json:"image"`
		WorkingDays    []string      `json:"working_days"`
		Tables         []TableCreate `json:"tables"`
		Locations      []Location    `json:"locations"`
	}

	RestaurantFilter struct {
		Name             string   `json:"name"`
		FoundationYear   string   `json:"foundation_year"`
		WorkingDays      []string `json:"working_days"`
		CustomerLocation Location `json:"customer_location"`
		SearchRadius     float64  `json:"search_radius"`
	}

	RestaurantSummary struct {
		ID          string     `json:"id"`
		Name        string     `json:"name"`
		PhoneNumber string     `json:"phone_number"`
		OpeningTime string     `json:"opening_time"`
		ClosingTime string     `json:"closing_time"`
		ImageURL    string     `json:"image_url"`
		WorkingDays []string   `json:"working_days"`
		Locations   []Location `json:"locations"`
	}

	Restaurant struct {
		ID             string     `json:"id"`
		Name           string     `json:"name"`
		FoundationYear string     `json:"foundation_year"`
		PhoneNumber    string     `json:"phone_number"`
		OpeningTime    string     `json:"opening_time"`
		ClosingTime    string     `json:"closing_time"`
		ImageURL       string     `json:"image_url"`
		WorkingDays    []string   `json:"working_days"`
		Tables         []Table    `json:"tables"`
		Locations      []Location `json:"locations"`
	}

	ImageURLCreatorFunc func(imageName valueobject.Image) string

	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
)

func (rc RestaurantCreate) ToAggregate() (*aggregate.Restaurant, error) {
	verifiedName, err := valueobject.NewRestaurantName(rc.Name)
	if err != nil {
		return nil, err
	}

	verifiedFoundationYear, err := valueobject.NewFoundationYear(rc.FoundationYear)
	if err != nil {
		return nil, err
	}

	verifiedPhoneNumber, err := valueobject.NewPhoneNumber(rc.PhoneNumber)
	if err != nil {
		return nil, err
	}

	verifiedOpeningTime, err := valueobject.NewWorkTime(rc.OpeningTime)
	if err != nil {
		return nil, err
	}

	verifiedClosingTime, err := valueobject.NewWorkTime(rc.ClosingTime)
	if err != nil {
		return nil, err
	}

	verifiedWorkingDays := make([]time.Weekday, 0)
	for _, workingDay := range rc.WorkingDays {
		verifiedWorkingDay, err := valueobject.ParseWeekday(workingDay)
		if err != nil {
			return nil, err
		}

		verifiedWorkingDays = append(verifiedWorkingDays, verifiedWorkingDay)
	}

	verifiedTables := make([]entity.Table, 0)
	for _, table := range rc.Tables {
		verifiedName, err := valueobject.NewTableName(table.Name)
		if err != nil {
			return nil, err
		}

		table := entity.NewTable()
		table.SetName(verifiedName)

		verifiedTables = append(verifiedTables, table)
	}

	verifiedLocations := make(valueobject.Locations, 0)
	for _, location := range rc.Locations {
		verifiedLocation, err := valueobject.NewLocation(location.Latitude, location.Longitude)
		if err != nil {
			return nil, err
		}

		verifiedLocations = append(verifiedLocations, verifiedLocation)
	}

	restaurant := aggregate.NewRestaurant()
	restaurant.SetName(verifiedName)
	restaurant.SetFoundationYear(verifiedFoundationYear)
	restaurant.SetPhoneNumber(verifiedPhoneNumber)
	restaurant.SetOpeningTime(verifiedOpeningTime)
	restaurant.SetClosingTime(verifiedClosingTime)
	restaurant.AddWorkingDays(verifiedWorkingDays...)
	restaurant.AddTables(verifiedTables...)
	restaurant.AddLocations(verifiedLocations...)

	return restaurant, nil
}

func (rf RestaurantFilter) ApplyFilter(restaurants []*aggregate.Restaurant) []*aggregate.Restaurant {
	filteredRestaurants := make([]*aggregate.Restaurant, 0)

	for _, restaurant := range restaurants {
		if !core.IsEmptyString(rf.Name) {
			if !strings.Contains(restaurant.Name().String(), rf.Name) {
				continue
			}
		}

		if !core.IsEmptyString(rf.FoundationYear) {
			if strings.Compare(rf.FoundationYear, restaurant.FoundationYear().String()) != 0 {
				continue
			}
		}

		if !isMatchingWorkingDays(restaurant.WorkingDays(), rf.WorkingDays) {
			continue
		}

		customerLocation, err := valueobject.NewLocation(rf.CustomerLocation.Latitude, rf.CustomerLocation.Longitude)
		if err != nil {
			return filteredRestaurants
		}

		if !restaurant.IsInRadius(customerLocation, rf.SearchRadius) {
			continue
		}

		filteredRestaurants = append(filteredRestaurants, restaurant)
	}

	return filteredRestaurants
}

func ToRestaurantSummary(restaurant *aggregate.Restaurant, imageConverter ImageURLCreatorFunc) RestaurantSummary {
	return RestaurantSummary{
		ID:          restaurant.ID().String(),
		Name:        restaurant.Name().String(),
		PhoneNumber: restaurant.PhoneNumber().String(),
		OpeningTime: restaurant.OpeningTime().String(),
		ClosingTime: restaurant.ClosingTime().String(),
		WorkingDays: toWorkingDays(restaurant.WorkingDays()),
		Locations:   toLocations(restaurant.Locations()),
		ImageURL:    imageConverter(restaurant.ImageName()),
	}
}

func ToRestaurantSummaryList(restaurants []*aggregate.Restaurant, imageConverter ImageURLCreatorFunc) []RestaurantSummary {
	rs := make([]RestaurantSummary, 0)

	for _, restaurant := range restaurants {
		rs = append(rs, ToRestaurantSummary(restaurant, imageConverter))
	}

	return rs
}

func ToRestaurant(restaurant *aggregate.Restaurant, imageConvertor ImageURLCreatorFunc) Restaurant {
	return Restaurant{
		ID:             restaurant.ID().String(),
		Name:           restaurant.Name().String(),
		FoundationYear: restaurant.FoundationYear().String(),
		PhoneNumber:    restaurant.PhoneNumber().String(),
		OpeningTime:    restaurant.OpeningTime().String(),
		ClosingTime:    restaurant.ClosingTime().String(),
		WorkingDays:    toWorkingDays(restaurant.WorkingDays()),
		ImageURL:       imageConvertor(restaurant.ImageName()),
		Tables:         toTableList(restaurant.Tables()),
		Locations:      toLocations(restaurant.Locations()),
	}
}

func ToRestaurantList(restaurants []*aggregate.Restaurant, imageConverter ImageURLCreatorFunc) []Restaurant {
	r := make([]Restaurant, 0)

	for _, restaurant := range restaurants {
		r = append(r, ToRestaurant(restaurant, imageConverter))
	}

	return r
}

func toWorkingDays(workingDays []time.Weekday) []string {
	wds := make([]string, 0)
	for _, wd := range workingDays {
		wds = append(wds, wd.String())
	}

	return wds
}

func isMatchingWorkingDays(src []time.Weekday, filter []string) bool {
	for _, workingDay := range filter {
		verifiedWeekday, err := valueobject.ParseWeekday(workingDay)
		if err != nil {
			return false
		}

		if !slices.Contains(src, verifiedWeekday) {
			return false
		}
	}

	return true
}

func toLocations(locations valueobject.Locations) []Location {
	ls := make([]Location, 0)
	for _, l := range locations {
		ls = append(ls, Location{
			Latitude:  l.Lat(),
			Longitude: l.Long(),
		})
	}

	return ls
}
