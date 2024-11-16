package valueobject

import (
	"errors"
	"fmt"
	"strings"
)

type Currency int64

type Price struct {
	quantity float64
	currency Currency
}

const (
	TRY Currency = iota
)

func ParseCurrency(currency string) (Currency, error) {
	currency = strings.TrimSpace(currency)
	currency = strings.ToLower(currency)

	switch currency {
	case "try":
		return TRY, nil
	default:
		return -1, errors.New("unsupported currency")
	}
}

func AvailableCurrencies() []string {
	return []string{
		TRY.String(),
	}
}

func (c Currency) String() string {
	switch c {
	case TRY:
		return "TRY"
	default:
		return ""
	}
}

func NewPrice(quantity float64, currency Currency) Price {
	return Price{
		quantity: quantity,
		currency: currency,
	}
}

func (p Price) Quantity() float64 {
	return p.quantity
}

func (p Price) Currency() Currency {
	return p.currency
}

func (p Price) String() string {
	return fmt.Sprintf("%.2f %s", p.quantity, p.currency)
}
