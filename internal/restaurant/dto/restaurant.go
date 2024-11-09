package dto

import (
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
	}

	RestaurantFilter struct {
		Name           string `json:"name"`
		FoundationYear string `json:"foundation_year"`
	}

	RestaurantSummary struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		PhoneNumber string   `json:"phone_number"`
		OpeningTime string   `json:"opening_time"`
		ClosingTime string   `json:"closing_time"`
		ImageURL    string   `json:"image_url"`
		WorkingDays []string `json:"working_days"`
	}

	Restaurant struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		FoundationYear string   `json:"foundation_year"`
		PhoneNumber    string   `json:"phone_number"`
		OpeningTime    string   `json:"opening_time"`
		ClosingTime    string   `json:"closing_time"`
		ImageURL       string   `json:"image_url"`
		WorkingDays    []string `json:"working_days"`
		Tables         []Table  `json:"tables"`
	}

	ImageURLCreatorFunc func(imageName valueobject.Image) string
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

	restaurant := aggregate.NewRestaurant()
	restaurant.SetName(verifiedName)
	restaurant.SetFoundationYear(verifiedFoundationYear)
	restaurant.SetPhoneNumber(verifiedPhoneNumber)
	restaurant.SetOpeningTime(verifiedOpeningTime)
	restaurant.SetClosingTime(verifiedClosingTime)
	restaurant.AddWorkingDays(verifiedWorkingDays...)
	restaurant.AddTables(verifiedTables...)

	return restaurant, nil
}

func (rf RestaurantFilter) ApplyFilter(restaurant *aggregate.Restaurant) bool {
	if !core.IsEmptyString(rf.Name) {
		if strings.Compare(rf.Name, restaurant.Name().String()) != 0 {
			return false
		}
	}

	if !core.IsEmptyString(rf.FoundationYear) {
		if strings.Compare(rf.FoundationYear, restaurant.FoundationYear().String()) != 0 {
			return false
		}
	}

	return true
}

func ToRestaurantSummary(restaurant *aggregate.Restaurant, imageConverter ImageURLCreatorFunc) RestaurantSummary {
	workingDays := make([]string, 0)
	for _, wd := range restaurant.WorkingDays() {
		workingDays = append(workingDays, wd.String())
	}

	return RestaurantSummary{
		ID:          restaurant.ID().String(),
		Name:        restaurant.Name().String(),
		PhoneNumber: restaurant.PhoneNumber().String(),
		OpeningTime: restaurant.OpeningTime().String(),
		ClosingTime: restaurant.ClosingTime().String(),
		WorkingDays: workingDays,
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
	workingDays := make([]string, 0)
	for _, wd := range restaurant.WorkingDays() {
		workingDays = append(workingDays, wd.String())
	}

	return Restaurant{
		ID:             restaurant.ID().String(),
		Name:           restaurant.Name().String(),
		FoundationYear: restaurant.FoundationYear().String(),
		PhoneNumber:    restaurant.PhoneNumber().String(),
		OpeningTime:    restaurant.OpeningTime().String(),
		ClosingTime:    restaurant.ClosingTime().String(),
		WorkingDays:    workingDays,
		ImageURL:       imageConvertor(restaurant.ImageName()),
		Tables:         ToTableList(restaurant.Tables()),
	}
}
