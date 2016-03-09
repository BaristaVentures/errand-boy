package repos

import (
	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/gorilla/mux"
)

// Router returns a preconfigured *mux.Router.
func Router(router *mux.Router) *mux.Router {
	routes := routers.Routes{
		&routers.Route{
			Path:    "/pr",
			Method:  "POST",
			Handler: pullRequestHandler,
		},
	}

	for _, r := range routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
	}
	return router
}
