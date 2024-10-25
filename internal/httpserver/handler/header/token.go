package header

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrInvalidBearerToken = errors.New("invalid bearer token format")
)

func GetBearerToken(header http.Header) (string, error) {
	tokenChunks := strings.Fields(header.Get(AuthorizationKey))

	if len(tokenChunks) < 2 {
		return "", ErrInvalidBearerToken
	}

	return tokenChunks[1], nil
}
