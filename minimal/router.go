package minimal

import (
	"fmt"
	"net/http"
)

// Register exposes endpoints to be accesible for the API
func Register(routes ...Route) {
	registerNotFound()
	for _, route := range routes {
		uri := fmt.Sprintf("/%s", route.URI)

		http.HandleFunc(uri, Common(route))
		// endpoint != endpoint/
		http.HandleFunc(uri+"/", Common(route))
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"error\": \"Not found\"}")
}

func registerNotFound() {
	http.HandleFunc("/", notFoundHandler)
}
