package entity

import (
	"slices"

	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/menu/domain/valueobject"
)

type MenuItem struct {
	core.Entity
	name               valueobject.MenuItemName
	description        string
	imageName          valueobject.Image
	price              valueobject.Price
	discountPercentage valueobject.Percentage
	ingredients        []string
}

func NewMenuItem() MenuItem {
	return MenuItem{
		Entity:      core.NewEntity(),
		ingredients: make([]string, 0),
	}
}

func (m *MenuItem) Name() valueobject.MenuItemName {
	return m.name
}

func (m *MenuItem) SetName(name valueobject.MenuItemName) {
	m.name = name
}

func (m *MenuItem) Description() string {
	return m.description
}

func (m *MenuItem) SetDescription(description string) {
	m.description = description
}

func (m *MenuItem) ImageName() valueobject.Image {
	return m.imageName
}

func (m *MenuItem) SetImageName(image valueobject.Image) {
	m.imageName = image
}

func (m *MenuItem) DiscountedPrice() valueobject.Price {
	discountAmount := (m.price.Amount() * float64(m.discountPercentage.Amount())) / 100
	discountedPrice := (m.price.Amount() - discountAmount)

	return valueobject.NewPrice(
		discountedPrice,
		m.price.Currency(),
	)
}

func (m *MenuItem) Price() valueobject.Price {
	return m.price
}

func (m *MenuItem) SetPrice(price valueobject.Price) {
	m.price = price
}

func (m *MenuItem) DiscountPercentage() valueobject.Percentage {
	return m.discountPercentage
}

func (m *MenuItem) SetDiscountPercentage(discountPercentage valueobject.Percentage) {
	m.discountPercentage = discountPercentage
}

func (m *MenuItem) Ingredients() []string {
	return m.ingredients
}

func (m *MenuItem) AddIngredients(ingredients ...string) {
	for _, i := range ingredients {
		if slices.Contains(m.ingredients, i) {
			continue
		}

		m.ingredients = append(m.ingredients, i)
	}
}
