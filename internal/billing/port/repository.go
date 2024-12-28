package port

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/billing/dto"
)

type (
	BillRepository interface {
		Save(bill *aggregate.Bill) error
		GetBySessionID(sessionID uuid.UUID) (*aggregate.Bill, error)
	}

	BillViewRepository interface {
		GetPostOrders(sessionID uuid.UUID) ([]dto.PostOrder, error)
		GetPastBills(customerID uuid.UUID) ([]dto.PastBill, error)
	}
)
