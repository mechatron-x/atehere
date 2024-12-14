package dto

import (
	"fmt"

	"github.com/google/uuid"
)

type (
	CheckoutEvent struct {
		ID           uuid.UUID
		RestaurantID string
		Table        string
		InvokeTime   int64
	}

	Session struct {
		SessionID string `json:"session_id"`
	}
)

func (scv CheckoutEvent) Message() string {
	return fmt.Sprintf("Table %s has requested checkout",
		scv.Table,
	)
}
