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

func NewPrice(amount float64, currency Currency) (Price, error) {
	if amount < 0 {
		return Price{}, errors.New("price cannot be smaller than 0")
	}
	return Price{
		amount:   amount,
		currency: currency,
	}, nil
}

func MustPrice(amount float64, currency Currency) Price {
	price, err := NewPrice(amount, currency)
	if err != nil {
		panic(err)
	}

	return price
}

func (p Price) Amount() float64 {
	return p.amount
}

func (p Price) Currency() Currency {
	return p.currency
}

func (p Price) IsZero() bool {
	return p.currency == 0
}

func (p Price) Equals(price Price) bool {
	return p.amount == price.amount
}

func (p Price) Add(price Price) Price {
	return MustPrice(p.amount+price.amount, price.currency)
}

func (p Price) Subtract(price Price) (Price, error) {
	remaining := p.amount - price.amount
	return NewPrice(remaining, price.currency)
}

func (p Price) Multiply(factor float64) Price {
	return MustPrice(p.amount*factor, p.currency)
}

func (p Price) String() string {
	return fmt.Sprintf("%.2f %s", p.amount, p.currency)
}
