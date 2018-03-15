package response

import (
	"net/http"
	"strings"
)

// EnableCORS support to the HTTP server by setting response headers
func EnableCORS(w http.ResponseWriter) {
	allowedHeaders := []string{
		"DNT",
		"User-Agent",
		"X-Requested-With",
		"If-Modified-Since",
		"Cache-Control",
		"Content-Type",
		"Range",
	}
	allowedMethods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
	}
	exposedHeaders := []string{
		"Content-Length",
		"Content-Range",
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", join(allowedMethods))
	w.Header().Set("Access-Control-Allow-Headers", join(allowedHeaders))
	w.Header().Set("Access-Control-Expose-Headers", join(exposedHeaders))
}

func join(slice []string) string {
	return strings.Join(slice, ",")
}
