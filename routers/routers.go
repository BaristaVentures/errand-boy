package routers

import "net/http"

// Route represents an endpoint route.
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes is a Route array.
type Routes []*Route
