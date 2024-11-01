package request

import (
	"encoding/json"
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

func Decode(r *http.Request, w http.ResponseWriter, to any) error {
	err := json.NewDecoder(r.Body).Decode(to)
	if err != nil {
		response.EncodeError(w, err, http.StatusBadRequest)
		return err
	}

	return nil
}
