package dto

type (
	PayBillItem struct {
		BillItemID string `json:"bill_item_id"`
		Quantity   int    `json:"quantity"`
	}

	PayBillItems struct {
		BillItems []PayBillItem `json:"bill_items"`
	}
)
