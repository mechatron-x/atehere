package service

type BillingService struct {
}

func NewBilling() *BillingService {
	return &BillingService{}
}

func (rcv *BillingService) CreateBill(billingMethod string) {

}
