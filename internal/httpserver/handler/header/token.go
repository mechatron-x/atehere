package header

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(header http.Header) (string, error) {
	tokenChunks := strings.Fields(header.Get(AuthorizationKey))

	if len(tokenChunks) < 2 {
		return "", errors.New("invalid bearer token format")
	}

	return tokenChunks[1], nil
}
