package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/restaurant/dto"
	"github.com/mechatron-x/atehere/internal/restaurant/port"
)

type Restaurant struct {
	restaurantRepo port.RestaurantRepository
	authService    port.Authenticator
	imageStorage   port.ImageStorage
	apiConf        config.Api
}

const (
	defaultPageSize int = 20
)

func NewRestaurant(
	restaurantRepo port.RestaurantRepository,
	authService port.Authenticator,
	fileService port.ImageStorage,
	apiConf config.Api,
) *Restaurant {
	return &Restaurant{
		restaurantRepo: restaurantRepo,
		authService:    authService,
		imageStorage:   fileService,
		apiConf:        apiConf,
	}
}

func (rs *Restaurant) Create(idToken string, createDto dto.RestaurantCreate) (*dto.Restaurant, error) {
	restaurant, err := createDto.ToAggregate()
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	managerID, err := rs.authService.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	imageName, err := rs.imageStorage.Save(restaurant.ID().String(), createDto.Image)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	verifiedImage, err := valueobject.NewImage(imageName)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	restaurant.SetOwner(uuid.MustParse(managerID))
	restaurant.SetImageName(verifiedImage)

	err = rs.restaurantRepo.Save(restaurant)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	restaurantDto := dto.ToRestaurant(restaurant, rs.createImageURL)
	return &restaurantDto, nil
}

func (rs *Restaurant) GetByID(id string) (*dto.RestaurantSummary, error) {
	verifiedID, err := uuid.Parse(id)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	restaurant, err := rs.restaurantRepo.GetByID(verifiedID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	summaryDto := dto.ToRestaurantSummary(restaurant, rs.createImageURL)
	return &summaryDto, nil
}

func (rs *Restaurant) List(page int, filterDto dto.RestaurantFilter) ([]dto.RestaurantSummary, error) {
	restaurants, err := rs.restaurantRepo.GetAll()
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	filteredRestaurants := filterDto.ApplyFilter(restaurants)

	return dto.ToRestaurantSummaryList(filteredRestaurants, rs.createImageURL), nil
}

func (rs *Restaurant) AvailableWorkingDays() []string {
	return valueobject.AvailableWeekdays()
}

func (rs *Restaurant) FoundationYearFormat() string {
	return valueobject.FoundationYearFormat
}

func (rs *Restaurant) WorkingTimeFormat() string {
	return valueobject.WorkingTimeFormat
}

func (rs *Restaurant) createImageURL(imageName valueobject.Image) string {
	if core.IsEmptyString(imageName.String()) {
		return ""
	}

	return fmt.Sprintf("%s/static/%s", rs.apiConf.URL, imageName.String())
}
