package dto

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	SessionClosedEvent struct {
		ID           uuid.UUID
		RestaurantID string
		Table        string
		InvokeTime   int64
	}
)

func (scv SessionClosedEvent) Message() string {
	return fmt.Sprintf("Session of table %s has been closed",
		scv.Table,
	)
}
