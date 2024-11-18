package service

import (
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

func (ms *Menu) Create(idToken string, createDto *dto.MenuCreate) (*dto.Menu, error) {
	menu, err := createDto.ToAggregate()
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	if err := ms.verifyOwnership(idToken, createDto.RestaurantID); err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	if err := ms.repository.Save(menu); err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return dto.ToMenu(menu, ms.createImageURL), nil
}

func (ms *Menu) AddMenuItem(idToken string, createDto *dto.MenuItemCreate) (*dto.Menu, error) {
	if err := ms.verifyOwnership(idToken, createDto.RestaurantID); err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	menuID, err := uuid.Parse(createDto.MenuID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	menuItem, err := createDto.ToEntity()
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	menu, err := ms.repository.GetByID(menuID)
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

// TODO: Add menu item delete method

func (ms *Menu) ListForCustomer(filterDto *dto.MenuFilter) ([]dto.Menu, error) {
	verifiedRestaurantID, err := uuid.Parse(filterDto.RestaurantID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	menus, err := ms.repository.GetManyByRestaurantID(verifiedRestaurantID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	return dto.ToMenuList(menus, ms.createImageURL), nil
}

func (ms *Menu) createImageURL(imageName valueobject.Image) string {
	if core.IsEmptyString(imageName.String()) {
		return ""
	}

	return fmt.Sprintf("%s/static/%s", ms.apiConf.URL, imageName.String())
}

func (ms *Menu) verifyOwnership(idToken, restaurantID string) error {
	verifiedRestaurantID, err := uuid.Parse(restaurantID)
	if err != nil {
		return err
	}

	managerID, err := ms.authService.GetUserID(idToken)
	if err != nil {
		return err
	}

	verifiedManagerID, err := uuid.Parse(managerID)
	if err != nil {
		return err
	}

	if !ms.repository.IsRestaurantOwner(verifiedRestaurantID, verifiedManagerID) {
		return err
	}

	return nil
}
