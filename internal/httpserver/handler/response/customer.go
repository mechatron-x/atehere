package response

import "github.com/mechatron-x/atehere/internal/usermanagement/dto"

type (
	CustomerSignUp struct {
		*dto.Customer
	}

	CustomerUpdateProfile struct {
		*dto.Customer
	}

	CustomerGetProfile struct {
		*dto.Customer
	}
)
