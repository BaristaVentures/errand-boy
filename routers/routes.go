package routers

import "github.com/plimble/ace"

// Route represents an endpoint route.
type Route struct {
	Path     string
	Method   string
	Handlers []ace.HandlerFunc
}

// Routes is a Route array.
type Routes []*Route

// Router identifies a struct as a router.
type Router interface {
	GetRoutes() Routes
}
