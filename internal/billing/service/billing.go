package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
	"github.com/mechatron-x/atehere/internal/billing/dto"
	"github.com/mechatron-x/atehere/internal/billing/port"
	"github.com/mechatron-x/atehere/internal/core"
	"github.com/mechatron-x/atehere/internal/infrastructure/logger"
	"go.uber.org/zap"
)

type BillingService struct {
	authenticator                 port.Authenticator
	repository                    port.BillRepository
	allPaymentsDoneEventPublisher port.AllPaymentsDoneEventPublisher
	log                           *zap.Logger
}

func NewBilling(
	authenticator port.Authenticator,
	repository port.BillRepository,
	allPaymentsDoneEventPublisher port.AllPaymentsDoneEventPublisher,
) *BillingService {
	return &BillingService{
		authenticator:                 authenticator,
		repository:                    repository,
		allPaymentsDoneEventPublisher: allPaymentsDoneEventPublisher,
		log:                           logger.Instance(),
	}
}

func (rcv *BillingService) Get(idToken string, sessionID string) (*dto.Bill, error) {
	requesterID, err := rcv.authenticator.GetUserID(idToken)
	if err != nil {
		return nil, core.NewUnauthorizedError(err)
	}

	verifiedRequesterID, err := uuid.Parse(requesterID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	verifiedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		return nil, core.NewValidationFailureError(err)
	}

	bill, err := rcv.repository.GetBySessionID(verifiedSessionID)
	if err != nil {
		return nil, core.NewResourceNotFoundError(err)
	}

	return dto.FromBill(verifiedRequesterID, bill), nil
}

func (rcv *BillingService) Pay(idToken string, sessionID string, billItems *dto.PayBillItems) error {
	paidBy, err := rcv.authenticator.GetUserID(idToken)
	if err != nil {
		return core.NewUnauthorizedError(err)
	}

	verifiedPaidBy, err := uuid.Parse(paidBy)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	verifiedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		return core.NewValidationFailureError(err)
	}

	bill, err := rcv.repository.GetBySessionID(verifiedSessionID)
	if err != nil {
		return core.NewResourceNotFoundError(err)
	}

	errs := make([]error, 0)
	for _, bi := range billItems.BillItems {
		verifiedBillItemID, err := uuid.Parse(bi.BillItemID)
		if err != nil {
			continue
		}

		verifiedCurrency, err := valueobject.ParseCurrency(bi.Currency)
		if err != nil {
			continue
		}

		verifiedPrice, err := valueobject.NewPrice(bi.Amount, verifiedCurrency)
		if err != nil {
			continue
		}

		err = bill.Pay(verifiedPaidBy, verifiedBillItemID, verifiedPrice)
		if err != nil {
			errs = append(errs, err)
		}
	}

	rcv.pushEventsAsync(bill.Events())
	if len(errs) != 0 {
		return core.NewDomainIntegrityViolationError(errors.Join(errs...))
	}

	return rcv.repository.Save(bill)
}

func (ss *BillingService) pushEventsAsync(events []core.DomainEvent) {
	go func(events []core.DomainEvent) {
		for _, e := range events {
			if allPaymentsDoneEvent, ok := e.(core.AllPaymentsDoneEvent); ok {
				ss.allPaymentsDoneEventPublisher.NotifyEvent(allPaymentsDoneEvent)
			} else {
				ss.log.Warn("unsupported event type skipping event processing")
			}
		}
	}(events)
}
