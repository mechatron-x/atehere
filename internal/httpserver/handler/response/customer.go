package response

import "github.com/mechatron-x/atehere/internal/usermanagement/dto"

type (
	SignUpCustomer struct {
		*dto.Customer
	}

	CustomerProfile struct {
		*dto.CustomerProfile
	}
)
