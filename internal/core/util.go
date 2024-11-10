package core

import (
	"strings"
)

func IsEmptyString(data string) bool {
	data = strings.TrimSpace(data)
	return len(data) == 0
}
