/*Todo:
 * Process env on port
 * HTTP verbs
 * Database integration
 * CORS
 * Middleware
 * Error handlers
 */
package main

import (
	"log"
	"net/http"

	"github.com/elbuki/minimal/router"
)

// Post struct that is used as a mock
type Post struct {
	ID    uint
	Title string
	Body  string
}

func homeHandler() interface{} {
	post := Post{
		ID:    1,
		Title: "Post title",
		Body:  "Post message that describes the post",
	}

	return post
}

func main() {
	routes := []router.Route{
		router.Route{URI: "", Handler: homeHandler},
	}

	router.Register(routes)

	log.Println("Server started at port 3000")
	http.ListenAndServe(":3000", nil)
}
