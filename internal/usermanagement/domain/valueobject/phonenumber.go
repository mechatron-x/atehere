package valueobject

type PhoneNumber struct {
	phoneNumber string
}

func NewPhoneNumber(phoneNumber string) (PhoneNumber, error) {
	return PhoneNumber{phoneNumber}, nil
}

func (p PhoneNumber) String() string {
	return p.phoneNumber
}
