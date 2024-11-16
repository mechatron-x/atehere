package dto

import (
	"github.com/mechatron-x/atehere/internal/menu/domain/entity"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
)

type (
	MenuItemCreate struct {
		RestaurantID       string   `json:"restaurant_id"`
		MenuID             string   `json:"menu_id"`
		Name               string   `json:"name"`
		Description        string   `json:"description"`
		Image              string   `json:"image"`
		Price              Price    `json:"price"`
		DiscountPercentage int      `json:"discount_percentage"`
		Ingredients        []string `json:"ingredients"`
	}

	MenuItem struct {
		ID                 string   `json:"id"`
		Name               string   `json:"name"`
		Description        string   `json:"description"`
		ImageURL           string   `json:"image_url"`
		Price              Price    `json:"price"`
		DiscountPercentage int      `json:"discount_percentage"`
		DiscountedPrice    Price    `json:"discounted_price"`
		Ingredients        []string `json:"ingredients"`
	}

	Price struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}

	ImageURLCreatorFunc func(imageName valueobject.Image) string
)

func (mic MenuItemCreate) ToEntity() (*entity.MenuItem, error) {
	verifiedName, err := valueobject.NewMenuItemName(mic.Name)
	if err != nil {
		return nil, err
	}

	verifiedDescription := mic.Description

	verifiedPriceQuantity := mic.Price.Amount
	verifiedPriceCurrency, err := valueobject.ParseCurrency(mic.Price.Currency)
	if err != nil {

	}
	verifiedPrice := valueobject.NewPrice(verifiedPriceQuantity, verifiedPriceCurrency)

	verifiedDiscountPercentage, err := valueobject.NewPercentage(mic.DiscountPercentage)
	if err != nil {
		return nil, err
	}

	verifiedIngredients := mic.Ingredients

	menuItem := entity.NewMenuItem()
	menuItem.SetName(verifiedName)
	menuItem.SetDescription(verifiedDescription)
	menuItem.SetPrice(verifiedPrice)
	menuItem.SetDiscountPercentage(verifiedDiscountPercentage)
	menuItem.AddIngredients(verifiedIngredients...)

	return &menuItem, nil
}

func toMenuItem(menuItem entity.MenuItem, imageConvertor ImageURLCreatorFunc) MenuItem {
	priceDto := Price{
		Amount:   menuItem.Price().Quantity(),
		Currency: menuItem.Price().Currency().String(),
	}

	discountedPriceDto := Price{
		Amount:   menuItem.DiscountedPrice().Quantity(),
		Currency: menuItem.DiscountedPrice().Currency().String(),
	}

	return MenuItem{
		ID:                 menuItem.ID().String(),
		Name:               menuItem.Name().String(),
		Description:        menuItem.Description(),
		ImageURL:           imageConvertor(menuItem.ImageName()),
		Price:              priceDto,
		DiscountPercentage: menuItem.DiscountPercentage().Amount(),
		DiscountedPrice:    discountedPriceDto,
		Ingredients:        menuItem.Ingredients(),
	}
}

func toMenuItemList(menuItems []entity.MenuItem, imageCreator ImageURLCreatorFunc) []MenuItem {
	menuItemDtos := make([]MenuItem, 0)
	for _, mi := range menuItems {
		menuItemDtos = append(menuItemDtos, toMenuItem(mi, imageCreator))
	}

	return menuItemDtos
}
