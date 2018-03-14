package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler from HTTP requests signature
type Handler func(w http.ResponseWriter, r *http.Request)

// Common acts as a middleware that affects to all endpoints
func Common(action string, handler func() interface{}) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != action {
			fmt.Fprint(w, "{\"error\": \"Not found\"}")
			return
		}

		// TODO: Set status code depending on the action

		json.NewEncoder(w).Encode(handler())
	}
}
