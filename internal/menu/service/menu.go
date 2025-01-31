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

type MenuService struct {
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
) *MenuService {
	return &MenuService{
		repository:   repository,
		authService:  authService,
		imageStorage: fileService,
		apiConf:      apiConf,
	}
}

func (ms *MenuService) Create(idToken string, createDto *dto.MenuCreate) (*dto.Menu, error) {
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

func (ms *MenuService) AddMenuItem(idToken string, createDto *dto.MenuItemCreate) (*dto.Menu, error) {
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

	if err := ms.verifyOwnership(idToken, menu.RestaurantID().String()); err != nil {
		return nil, core.NewUnauthorizedError(err)
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
	err = menu.AddMenuItems(*menuItem)
	if err != nil {
		return nil, core.NewDomainIntegrityViolationError(err)
	}

	if err := ms.repository.Save(menu); err != nil {
		return nil, core.NewPersistenceFailureError(err)
	}

	return dto.ToMenu(menu, ms.createImageURL), nil
}

func (ms *MenuService) ListForCustomer(filterDto *dto.MenuFilter) ([]dto.Menu, error) {
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

func (ms *MenuService) createImageURL(imageName valueobject.Image) string {
	if core.IsEmptyString(imageName.String()) {
		return ""
	}

	return fmt.Sprintf("%s/static/%s", ms.apiConf.URL, imageName.String())
}

func (ms *MenuService) verifyOwnership(idToken, restaurantID string) error {
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
		return errors.New("invalid restaurant ownership")
	}

	return nil
}
