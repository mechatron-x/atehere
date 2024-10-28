package request

import "github.com/mechatron-x/atehere/internal/usermanagement/dto"

type (
	CustomerSignUp struct {
		dto.CustomerSignUp
	}

	CustomerUpdateProfile struct {
		dto.Customer
	}
)
