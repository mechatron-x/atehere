package dto

import (
	"github.com/google/uuid"
	"github.com/mechatron-x/atehere/internal/billing/domain/aggregate"
	"github.com/mechatron-x/atehere/internal/billing/domain/entity"
	"github.com/mechatron-x/atehere/internal/billing/domain/valueobject"
)

type (
	Bill struct {
		BillItems              []BillItem `json:"bill_items"`
		TotalDue               float64    `json:"total_due"`
		RemainingPrice         float64    `json:"remaining_price"`
		IndividualPaymentTotal float64    `json:"individual_payment_total"`
		Currency               string     `json:"currency"`
	}

	BillItem struct {
		ID                string  `json:"id"`
		OwnerID           string  `json:"owner_id"`
		ItemName          string  `json:"item_name"`
		UnitPrice         float64 `json:"unit_price"`
		Quantity          int     `json:"quantity"`
		TotalDue          float64 `json:"total_due"`
		RemainingPrice    float64 `json:"remaining_price"`
		IndividualPayment float64 `json:"individual_payment"`
		Currency          string  `json:"currency"`
		IsAllPaid         bool    `json:"is_all_paid"`
	}

	PayBillItem struct {
		BillItemID string  `json:"bill_item_id"`
		Amount     float64 `json:"amount"`
		Currency   string  `json:"currency"`
	}

	PayBillItems struct {
		BillItems []PayBillItem `json:"bill_items"`
	}
)

func FromBill(requesterID uuid.UUID, bill *aggregate.Bill) *Bill {
	billDTO := &Bill{
		BillItems: FromBillItems(requesterID, bill.BillItems()),
	}

	totalDue := 0.0
	remainingPrice := 0.0
	individualPaymentTotal := 0.0
	currency := ""

	for _, billItem := range billDTO.BillItems {
		totalDue += billItem.TotalDue
		remainingPrice += billItem.RemainingPrice
		individualPaymentTotal += billItem.IndividualPayment
		currency = billItem.Currency
	}

	billDTO.TotalDue = totalDue
	billDTO.RemainingPrice = remainingPrice
	billDTO.IndividualPaymentTotal = individualPaymentTotal
	billDTO.Currency = currency

	return billDTO
}

func FromBillItems(requesterID uuid.UUID, billItems []entity.BillItem) []BillItem {
	billItemDTOs := make([]BillItem, 0)

	for _, bi := range billItems {
		billItemDTOs = append(billItemDTOs, FromBillItem(requesterID, bi))
	}

	return billItemDTOs
}

func FromBillItem(requesterID uuid.UUID, billItem entity.BillItem) BillItem {
	individualPayment, ok := billItem.Payments()[requesterID]
	if !ok {
		individualPayment = valueobject.MustPrice(0, billItem.UnitPrice().Currency())
	}

	return BillItem{
		ID:                billItem.ID().String(),
		OwnerID:           billItem.OwnerID().String(),
		ItemName:          billItem.ItemName(),
		UnitPrice:         billItem.UnitPrice().Amount(),
		Quantity:          billItem.Quantity().Int(),
		TotalDue:          billItem.TotalDue().Amount(),
		RemainingPrice:    billItem.RemainingAmount().Amount(),
		Currency:          billItem.UnitPrice().Currency().String(),
		IndividualPayment: individualPayment.Amount(),
		IsAllPaid:         billItem.TotalDue().Equals(billItem.PaidAmount()),
	}
}
