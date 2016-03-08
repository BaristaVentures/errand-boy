package github

import (
	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/gorilla/mux"
)

var instance GitHubRouter

// GitHubRouter represents a GitHub endpoint router.
type GitHubRouter struct {
	routes routers.Routes
}

func init() {
	instance = GitHubRouter{}
	instance.routes = routers.Routes{
		&routers.Route{
			Path:    "/pr",
			Method:  "POST",
			Handler: pullRequestHandler,
		},
	}
}

// Instance returns this router's instance.
func Instance() routers.Router {
	return &instance
}

// SetUpRoutes sets up this router's routes.
func (gh *GitHubRouter) SetUpRoutes(router *mux.Router) {
	for _, r := range gh.routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
		// router.Handle(r.Method, r.Path, r.Handlers)
	}
}
