package routers

import "github.com/plimble/ace"

// GitHubRouter represents a GitHub endpoint router.
type GitHubRouter struct {
	routes Routes
}

// NewGitHubRouter initializes the GitHub routes.
func NewGitHubRouter() Router {
	gitHubRoutes := Routes{
		&Route{
			Path:     "/github/pr",
			Method:   "GET",
			Handlers: []ace.HandlerFunc{pullRequestHandler},
		},
	}
	return &GitHubRouter{routes: gitHubRoutes}
}

// GetRoutes returns a GitHub router's routes.
func (gh *GitHubRouter) GetRoutes() Routes {
	return gh.routes
}

func pullRequestHandler(c *ace.C) {
	c.String(200, "%s\n", "Errand Boy is running!")
}
