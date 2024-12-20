package valueobject

import (
	"errors"
	"fmt"
	"strings"
)

type Currency int64

type Price struct {
	amount   float64
	currency Currency
}

const (
	UNKNOWN Currency = iota
	TRY
	USD
)

func ParseCurrency(currency string) (Currency, error) {
	currency = strings.TrimSpace(currency)
	currency = strings.ToLower(currency)

	switch currency {
	case "try":
		return TRY, nil
	case "usd":
		return USD, nil
	default:
		return -1, errors.New("unsupported currency")
	}
}

func AvailableCurrencies() []string {
	return []string{
		TRY.String(),
		USD.String(),
	}
}

func (c Currency) String() string {
	switch c {
	case TRY:
		return "TRY"
	case USD:
		return "USD"
	default:
		return ""
	}
}

func NewPrice(amount float64, currency Currency) Price {
	return Price{
		amount:   amount,
		currency: currency,
	}
}

func (p Price) Amount() float64 {
	return p.amount
}

func (p Price) Currency() Currency {
	return p.currency
}

func (p Price) String() string {
	return fmt.Sprintf("%.2f %s", p.amount, p.currency)
}
