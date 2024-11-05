package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/restaurant/dto"
	"github.com/mechatron-x/atehere/internal/restaurant/port"
)

type Restaurant struct {
	restaurantRepo port.RestaurantRepository
	authService    port.Authenticator
}

func NewRestaurant(
	restaurantRepo port.RestaurantRepository,
	authService port.Authenticator,
) *Restaurant {
	return &Restaurant{
		restaurantRepo: restaurantRepo,
		authService:    authService,
	}
}

func (rs *Restaurant) Create(idToken string, createDto dto.RestaurantCreate) (*dto.Restaurant, error) {
	restaurant, err := rs.validateCreateDto(createDto)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	managerID, err := rs.authService.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	restaurant.SetOwner(uuid.MustParse(managerID))

	err = rs.restaurantRepo.Save(restaurant)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return rs.toDto(restaurant), nil
}

func (rs *Restaurant) validateCreateDto(createDto dto.RestaurantCreate) (*aggregate.Restaurant, error) {
	verifiedName, err := valueobject.NewRestaurantName(createDto.Name)
	if err != nil {
		return nil, err
	}

	verifiedFoundationYear, err := valueobject.NewFoundationYear(createDto.FoundationYear)
	if err != nil {
		return nil, err
	}

	verifiedPhoneNumber, err := valueobject.NewPhoneNumber(createDto.PhoneNumber)
	if err != nil {
		return nil, err
	}

	verifiedOpeningTime, err := valueobject.NewWorkTime(createDto.OpeningTime)
	if err != nil {
		return nil, err
	}
	verifiedClosingTime, err := valueobject.NewWorkTime(createDto.ClosingTime)
	if err != nil {
		return nil, err
	}

	verifiedWorkingDays := make([]time.Weekday, 0)

	if len(createDto.WorkingDays) == 0 {
		return nil, errors.New("working days cannot be 0")
	}

	for _, workingDay := range createDto.WorkingDays {
		verifiedWorkingDay, err := valueobject.ParseWeekday(workingDay)
		if err != nil {
			return nil, err
		}

		verifiedWorkingDays = append(verifiedWorkingDays, verifiedWorkingDay)
	}

	restaurant := aggregate.NewRestaurant()
	restaurant.SetName(verifiedName)
	restaurant.SetFoundationYear(verifiedFoundationYear)
	restaurant.SetPhoneNumber(verifiedPhoneNumber)
	restaurant.SetOpeningTime(verifiedOpeningTime)
	restaurant.SetClosingTime(verifiedClosingTime)
	restaurant.AddWorkingDays(verifiedWorkingDays...)

	return restaurant, nil
}

func (rs *Restaurant) toDto(restaurant *aggregate.Restaurant) *dto.Restaurant {
	workingDays := make([]string, 0)
	for _, wd := range restaurant.WorkingDays() {
		workingDays = append(workingDays, wd.String())
	}

	return &dto.Restaurant{
		ID:             restaurant.ID().String(),
		OwnerID:        restaurant.OwnerID().String(),
		Name:           restaurant.Name().String(),
		FoundationYear: restaurant.FoundationYear().String(),
		PhoneNumber:    restaurant.PhoneNumber().String(),
		OpeningTime:    restaurant.OpeningTime().String(),
		ClosingTime:    restaurant.ClosingTime().String(),
		WorkingDays:    workingDays,
	}
}
