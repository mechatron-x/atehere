package response

import "github.com/mechatron-x/atehere/internal/usermanagement/dto"

type (
	ManagerSignUp struct {
		*dto.Manager
	}

	ManagerUpdateProfile struct {
		*dto.Manager
	}

	ManagerGetProfile struct {
		*dto.Manager
	}
)
