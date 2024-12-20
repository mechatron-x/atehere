package broker

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	ID() uuid.UUID
	InvokeTime() time.Time
}
