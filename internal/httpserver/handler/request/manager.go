package request

import "github.com/mechatron-x/atehere/internal/usermanagement/dto"

type (
	ManagerSignUp struct {
		dto.ManagerSignUp
	}

	ManagerUpdateProfile struct {
		dto.Manager
	}
)
