package response

import (
	"encoding/json"
	"net/http"
)

func Encode(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
}
