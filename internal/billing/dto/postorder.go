package dto

import (
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
)

type PostOrder struct {
	SessionID    string
	CustomerID   string
	MenuItemName string
	Quantity     int
	UnitPrice    float64
	Currency     string
}

func (o PostOrder) ToBillItem() (*entity.BillItem, error) {
	billItemBuilder := entity.NewBillItemBuilder()

	verifiedCurrency, err := valueobject.ParseCurrency(o.Currency)
	if err != nil {
		return nil, err
	}
	price := valueobject.NewPrice(o.UnitPrice, verifiedCurrency)

	verifiedQuantity, err := valueobject.NewQuantity(o.Quantity)
	if err != nil {
		return nil, err
	}

	return billItemBuilder.
		SetOwnerID(o.CustomerID).
		SetItemName(o.MenuItemName).
		SetPrice(price).
		SetQuantity(verifiedQuantity).
		Build()
}
