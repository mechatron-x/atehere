package request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mechatron-x/atehere/internal/httpserver/handler/response"
)

func Decode(r *http.Request, w http.ResponseWriter, to any) error {
	err := json.NewDecoder(r.Body).Decode(to)
	if err != nil {
		response.EncodeError(
			w,
			fmt.Errorf("failed to decode request body: %v", err), http.StatusBadRequest,
		)
		return err
	}

	return nil
}
