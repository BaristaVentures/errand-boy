package github

import (
	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/plimble/ace"
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
			Path:     "/pr",
			Method:   "POST",
			Handlers: []ace.HandlerFunc{pullRequestHandler},
		},
	}
}

// Instance returns this router's instance.
func Instance() routers.Router {
	return &instance
}

// SetUpRoutes sets up this router's routes.
func (gh *GitHubRouter) SetUpRoutes(router *ace.Router) {
	router = router.Group("/gh")
	for _, r := range gh.routes {
		router.Handle(r.Method, r.Path, r.Handlers)
	}
}
