/*Todo:
 * CORS
 * Middleware
 * Error handlers
 * Logging
 */
package main

import (
	"github.com/elbuki/minimal/router"
	"github.com/elbuki/minimal/server"
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

func addPost() interface{} {
	post := Post{
		ID:    2,
		Title: "New post title",
		Body:  "Post message that describes the post",
	}

	return post
}

func main() {
	routes := []router.Route{
		router.Route{URI: "posts", Action: "GET", Handler: homeHandler},
		router.Route{URI: "posts/create", Action: "POST", Handler: addPost},
	}

	router.Register(routes...)

	// Could get from environment variables
	server.Start(8080)
}
