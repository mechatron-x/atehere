package dto

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	SessionClosedEventView struct {
		ID           uuid.UUID
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
