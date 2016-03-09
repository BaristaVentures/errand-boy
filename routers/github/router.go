package github

import (
	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/gorilla/mux"
)

// Router returns a preconfigured *mux.Router.
func Router() *mux.Router {
	routes := routers.Routes{
		&routers.Route{
			Path:    "/pr",
			Method:  "POST",
			Handler: pullRequestHandler,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
	}
	return router
}
