package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler from HTTP requests signature
type Handler func(w http.ResponseWriter, r *http.Request)

// Route struct for HTTP handlers
type Route struct {
	URI     string
	handler func() string
}

func homeHandler() string {
	return "{\"asdf\": 5}"
}

func common(handler func() string) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, handler())
	}
}

func main() {
	var routes = []Route{
		Route{URI: "home", handler: homeHandler},
	}

	for _, route := range routes {
		uri := fmt.Sprintf("/%s", route.URI)

		http.HandleFunc(uri, common(route.handler))
	}

	log.Println("Server started at port 3000")
	http.ListenAndServe(":3000", nil)
}
