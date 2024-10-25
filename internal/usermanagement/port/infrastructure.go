package port

import (
	"github.com/mechatron-x/atehere/internal/core"
)

type (
	Authenticator interface {
		CreateUser(id, email, password string) core.PortError
		RevokeRefreshTokens(idToken string) core.PortError
		GetUserID(idToken string) (string, core.PortError)
		GetUserEmail(idToken string) (string, core.PortError)
	}
)
