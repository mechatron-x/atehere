package request

import (
	"encoding/json"
	"net/http"
)

func Decode(r *http.Request, w http.ResponseWriter, to any) error {
	err := json.NewDecoder(r.Body).Decode(to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	return nil
}
