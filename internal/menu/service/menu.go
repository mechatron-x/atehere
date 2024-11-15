package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/config"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/menu/dto"
	"github.com/mechatron-x/atehere/internal/menu/port"
)

type Menu struct {
	repository   port.MenuRepository
	authService  port.Authenticator
	imageStorage port.ImageStorage
	apiConf      config.Api
}

func NewMenu(
	repository port.MenuRepository,
	authService port.Authenticator,
	fileService port.ImageStorage,
	apiConf config.Api,
) *Menu {
	return &Menu{
		repository:   repository,
		authService:  authService,
		imageStorage: fileService,
		apiConf:      apiConf,
	}
}

func (ms *Menu) Create(idToken string, createDto dto.MenuCreate) (*dto.Menu, error) {
	menu, err := createDto.ToAggregate()
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	managerID, err := ms.authService.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}
	verifiedManagerID := uuid.MustParse(managerID)

	if !ms.repository.IsRestaurantOwner(menu.RestaurantID(), verifiedManagerID) {
		return nil, core.NewUnauthorizedError(errors.New("manager does not own the restaurant"))
	}

	if err := ms.repository.Save(menu); err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return dto.ToMenu(menu, ms.createImageURL), nil
}

func (ms *Menu) AddMenuItem(idToken string, createDto dto.MenuItemCreate) (*dto.Menu, error) {
	verifiedRestaurantID, err := uuid.Parse(createDto.RestaurantID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	managerID, err := ms.authService.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}
	verifiedManagerID := uuid.MustParse(managerID)

	if !ms.repository.IsRestaurantOwner(verifiedRestaurantID, verifiedManagerID) {
		return nil, core.NewUnauthorizedError(errors.New("manager does not own the restaurant"))
	}

	verifiedMenuID, err := uuid.Parse(createDto.MenuID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	menuItem, err := createDto.ToEntity()
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	menu, err := ms.repository.GetByID(verifiedMenuID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	imageName, err := ms.imageStorage.Save(menuItem.ID().String(), createDto.Image)
	if err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	verifiedImage, err := valueobject.NewImage(imageName)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	menuItem.SetImageName(verifiedImage)

	menu.AddMenuItems(*menuItem)
	if err := ms.repository.Save(menu); err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return dto.ToMenu(menu, ms.createImageURL), nil
}

func (ms *Menu) GetMenuByCategory(idToken, restaurantID, category string) (*dto.Menu, error) {
	return nil, nil
}

func (ms *Menu) createImageURL(imageName valueobject.Image) string {
	if core.IsEmptyString(imageName.String()) {
		return ""
	}

	return fmt.Sprintf("%s/static/%s", ms.apiConf.URL, imageName.String())
}
