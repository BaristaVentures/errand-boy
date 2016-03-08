package github

import (
	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/gorilla/mux"
)

var instance Router

// Router represents a GitHub endpoint router.
type Router struct {
	routes routers.Routes
}

// NewRouter eturns a new router with default routes configured.
func NewRouter() routers.Router {
	instance = Router{}
	instance.routes = routers.Routes{
		&routers.Route{
			Path:    "/pr",
			Method:  "POST",
			Handler: pullRequestHandler,
		},
	}
	return &instance
}

// SetUpRoutes sets up this router's routes.
func (gh *Router) SetUpRoutes(router *mux.Router) {
	for _, r := range gh.routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
	}
}
