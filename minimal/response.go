package minimal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler from HTTP requests signature
type Handler func(w http.ResponseWriter, r *http.Request)

// Common acts as a middleware that affects to all endpoints
func Common(route Route) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range route.Before {
			_, next, _ := middleware()

			if !next {
				return
			}
		}

		sendResponse(w, r, route)
	}
}

func sendResponse(w http.ResponseWriter, r *http.Request, route Route) {
	EnableCORS(w)
	w.Header().Set("Content-Type", "application/json")

	if r.Method != route.Action {
		fmt.Fprint(w, "{\"error\": \"Not found\"}")
		return
	}

	if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(route.Handler())
}
