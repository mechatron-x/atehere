package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode(r *http.Request, w http.ResponseWriter, to any) error {
	err := json.NewDecoder(r.Body).Decode(to)
	if err != nil {
		return fmt.Errorf("failed to decode request body: %v", err)
	}

	return nil
}
