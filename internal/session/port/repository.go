package port

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/session/domain/aggregate"
)

type SessionRepository interface {
	Save(session *aggregate.Session) error
	GetByTableID(tableID uuid.UUID) (*aggregate.Session, error)
	HasActiveSessions(tableID uuid.UUID) bool
}
