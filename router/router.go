package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler from HTTP requests signature
type handler func(w http.ResponseWriter, r *http.Request)

// Route struct for HTTP handlers
type Route struct {
	URI     string
	Action  string
	Handler func() interface{}
}

// Register exposes endpoints to be accesible for the API
func Register(routes ...Route) {
	for _, route := range routes {
		uri := fmt.Sprintf("/%s", route.URI)

		http.HandleFunc(uri, common(route.Action, route.Handler))
		// endpoint != endpoint/
		http.HandleFunc(uri+"/", common(route.Action, route.Handler))
	}
}

func common(action string, handler func() interface{}) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != action {
			fmt.Fprint(w, "{\"error\": \"Not found\"}")
			return
		}

		json.NewEncoder(w).Encode(handler())
	}
}
