package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents an endpoint route.
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
	Middleware
}

type Middleware func(http.Handler) http.Handler

// Routes is a Route array.
type Routes []*Route

// Router Identifies a struct as a router.
type Router interface {
	// Takes a *mux.Router to configure its internal routes.
	SetUpRoutes(*mux.Router)
}
