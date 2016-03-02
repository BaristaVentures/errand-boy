package routers

import (
	"testing"

	"github.com/hooklift/assert"
)

func TestNewGitHubRouter(t *testing.T) {
	ghRouter := NewGitHubRouter()
	assert.Cond(t, len(ghRouter.GetRoutes()) > 0, "The GitHub routes should not be empty.")
	for _, route := range ghRouter.GetRoutes() {
		assert.Cond(t, len(route.Handlers) > 0, "A route should always have at least one handler.")
		// TODO: check that the method is a valid REST one.
		assert.Cond(t, len(route.Method) > 0, "Invalid route method.")
		assert.Cond(t, len(route.Path) > 0, "Invalid route path.")
	}
}
