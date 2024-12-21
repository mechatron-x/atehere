package valueobject

import "strings"

type SessionState int64

const (
	Active SessionState = iota
	CheckoutPending
	Completed
)

func ParseSessionState(state string) SessionState {
	state = strings.TrimSpace(state)
	state = strings.ToLower(state)

	switch state {
	case "active":
		return Active
	case "checkout_pending", "checkoutPending":
		return CheckoutPending
	case "completed":
		return Completed
	default:
		return Active
	}
}

func (rcv SessionState) AvailableStates() []string {
	return []string{
		Active.String(),
		CheckoutPending.String(),
		Completed.String(),
	}
}

func (rcv SessionState) String() string {
	switch rcv {
	case Active:
		return "ACTIVE"
	case CheckoutPending:
		return "CHECKOUT_PENDING"
	case Completed:
		return "COMPLETED"
	default:
		return "ACTIVE"
	}
}
