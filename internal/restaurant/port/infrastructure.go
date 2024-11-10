package port

type (
	Authenticator interface {
		GetUserID(idToken string) (string, error)
		GetUserEmail(idToken string) (string, error)
	}

	ImageStorage interface {
		Save(fileName, data string) (string, error)
	}
)
