package main

import (
	"fmt"

	"github.com/elbuki/minimal/minimal"
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

func userExists() (interface{}, bool, int) {
	fmt.Println("userExists")
	return nil, true, 200
}

func isLoggedIn() (interface{}, bool, int) {
	fmt.Println("isLoggedIn")
	return "{}", false, 401
}

func main() {
	common := []func() (interface{}, bool, int){isLoggedIn, userExists}
	routes := []minimal.Route{
		minimal.Route{
			URI:     "posts",
			Action:  "GET",
			Handler: homeHandler,
			Before:  common,
		},
		minimal.Route{URI: "posts/create", Action: "POST", Handler: addPost},
	}

	minimal.Register(routes...)

	// Could get from environment variables
	minimal.Start(8080)
}
