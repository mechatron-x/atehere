package port

type (
	Authenticator interface {
		CreateUser(id, email, password string) error
		RevokeRefreshTokens(idToken string) error
		GetUserID(idToken string) (string, error)
		GetUserEmail(idToken string) (string, error)
	}
)
