package router

import (
	"fmt"
	"net/http"

	"github.com/elbuki/minimal/response"
)

// Route struct for HTTP handlers
type Route struct {
	URI     string
	Action  string
	Handler func() interface{}
}

// Register exposes endpoints to be accesible for the API
func Register(routes ...Route) {

	registerNotFound()
	for _, route := range routes {
		uri := fmt.Sprintf("/%s", route.URI)

		http.HandleFunc(uri, response.Common(route.Action, route.Handler))
		// endpoint != endpoint/
		http.HandleFunc(uri+"/", response.Common(route.Action, route.Handler))
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"error\": \"Not found\"}")
}

func registerNotFound() {
	http.HandleFunc("/", notFoundHandler)
}
