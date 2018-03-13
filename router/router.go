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
	Handler func() interface{}
}

// Register exposes endpoints to be accesible for the API
func Register(routes []Route) {
	for _, route := range routes {
		uri := fmt.Sprintf("/%s", route.URI)

		http.HandleFunc(uri, common(route.Handler))
	}
}

func common(handler func() interface{}) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(handler())
	}
}
