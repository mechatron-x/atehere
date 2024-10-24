package port

type (
	AuthRecord struct {
		UID           string
		Disabled      bool
		EmailVerified bool
		DisplayName   string
		Email         string
		PhoneNumber   string
	}

	Authenticator interface {
		CreateUser(id, email, password string) error
		RevokeRefreshTokens(idToken string) error
		VerifyUser(idToken string) (*AuthRecord, error)
	}
)
