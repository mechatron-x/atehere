package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/restaurant/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/restaurant/dto"
	"github.com/mechatron-x/atehere/internal/restaurant/port"
)

type RestaurantService struct {
	repository    port.RestaurantRepository
	authenticator port.Authenticator
	imageStorage  port.ImageStorage
	apiConf       config.Api
}

func NewRestaurant(
	repository port.RestaurantRepository,
	authenticator port.Authenticator,
	fileService port.ImageStorage,
	apiConf config.Api,
) *RestaurantService {
	return &RestaurantService{
		repository:    repository,
		authenticator: authenticator,
		imageStorage:  fileService,
		apiConf:       apiConf,
	}
}

func (rs *RestaurantService) Create(idToken string, createDto *dto.RestaurantCreate) (*dto.Restaurant, error) {
	restaurant, err := createDto.ToAggregate()
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	managerID, err := rs.authenticator.GetUserID(idToken)
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

	verifiedManagerID, err := uuid.Parse(managerID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	restaurant.SetOwner(verifiedManagerID)
	restaurant.SetImageName(verifiedImage)

	err = rs.repository.Save(restaurant)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	restaurantDto := dto.ToRestaurant(restaurant, rs.createImageURL)
	return &restaurantDto, nil
}

func (rs *RestaurantService) GetOneForCustomer(id string) (*dto.RestaurantSummary, error) {
	verifiedID, err := uuid.Parse(id)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	restaurant, err := rs.repository.GetByID(verifiedID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	summaryDto := dto.ToRestaurantSummary(restaurant, rs.createImageURL)
	return &summaryDto, nil
}

func (rs *RestaurantService) ListForManager(idToken string) ([]dto.Restaurant, error) {
	managerID, err := rs.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	verifiedManagerID, err := uuid.Parse(managerID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	restaurants, err := rs.repository.GetByOwnerID(verifiedManagerID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	return dto.ToRestaurantList(restaurants, rs.createImageURL), nil
}

func (rs *RestaurantService) ListForCustomer(filterDto *dto.RestaurantFilter) ([]dto.RestaurantSummary, error) {
	restaurants, err := rs.repository.GetAll()
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	filteredRestaurants := filterDto.ApplyFilter(restaurants)

	return dto.ToRestaurantSummaryList(filteredRestaurants, rs.createImageURL), nil
}

func (rs RestaurantService) Delete(idToken, restaurantID string) error {
	managerID, err := rs.authenticator.GetUserID(idToken)
	if err != nil {
		return core.NewUnauthorizedError(err)
	}

	verifiedManagerID, err := uuid.Parse(managerID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	verifiedRestaurantID, err := uuid.Parse(restaurantID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	restaurant, err := rs.repository.GetByID(verifiedRestaurantID)
	if err != nil {
		return core.NewResourceNotFoundError(err)
	}

	if !restaurant.IsOwner(verifiedManagerID) {
		return core.NewUnauthorizedError(errors.New("insufficient permissions to delete restaurant"))
	}

	restaurant.DeleteNow()

	return rs.repository.Save(restaurant)
}

func (rs *RestaurantService) AvailableWorkingDays() []string {
	return valueobject.AvailableWeekdays()
}

func (rs *RestaurantService) FoundationYearFormat() string {
	return valueobject.FoundationYearFormat
}

func (rs *RestaurantService) WorkingTimeFormat() string {
	return valueobject.WorkingTimeFormat
}

func (rs *RestaurantService) createImageURL(imageName valueobject.Image) string {
	if core.IsEmptyString(imageName.String()) {
		return ""
	}

	return fmt.Sprintf("%s/static/%s", rs.apiConf.URL, imageName.String())
}
