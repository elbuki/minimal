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
		EnableCORS(w)
		w.Header().Set("Content-Type", "application/json")

		if r.Method != action {
			fmt.Fprint(w, "{\"error\": \"Not found\"}")
			return
		}

		if r.Method == "POST" {
			w.WriteHeader(http.StatusCreated)
		}

		json.NewEncoder(w).Encode(handler())
	}
}
