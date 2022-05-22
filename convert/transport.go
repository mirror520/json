package convert

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/mirror520/json/endpoint"
)

func SnakeCaseHandler(endpoint endpoint.Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "404 page not found")
			return
		}

		var request map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err.Error())
			return
		}

		response, err := endpoint(context.Background(), request)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintln(w, err.Error())
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
