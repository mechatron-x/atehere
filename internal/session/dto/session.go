package dto

import "fmt"

type (
	SessionClosedEventView struct {
		RestaurantID string
		Table        string
		InvokeTime   int64
	}
)

func (scv SessionClosedEventView) Message() string {
	return fmt.Sprintf("Session of table %s has been closed",
		scv.Table,
	)
}
