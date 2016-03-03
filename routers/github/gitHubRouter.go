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

type pullRequestPayload struct {
	Action      string       `json:"action"`
	Number      int          `json:"number"`
	PullRequest *pullRequest `json:"pull_request"`
}

type pullRequest struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	MergedAt string `json:"merged_at"`
}

func init() {
	instance = GitHubRouter{}
	instance.routes = routers.Routes{
		&routers.Route{
			Path:     "/github/pr",
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
	for _, r := range gh.routes {
		router.Handle(r.Method, r.Path, r.Handlers)
	}
}
