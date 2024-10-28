package request

import "github.com/mechatron-x/atehere/internal/usermanagement/dto"

type (
	SignUpCustomer struct {
		dto.Customer
	}

	UpdateCustomer struct {
		dto.Customer
	}
)
