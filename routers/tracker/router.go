package tracker

import (
	"encoding/json"
	"net/http"

	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

// Route returns a configured *mux.Router.
func Route(router *mux.Router) *mux.Router {
	routes := &routers.Routes{
		&routers.Route{
			Path:    "/activity",
			Method:  "POST",
			Handler: activityHandler,
		},
	}
	for _, r := range *routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
	}

	return router
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	activity := &ActivityPayload{}
	json.NewDecoder(r.Body).Decode(&activity)
	// TODO: How to establish a relationship between te accepted story and the PR where the comment
	// should be made? Obvious option is getting the list of open PRs and checking which one has the
	// story's id in its title. Cons: the roundtrip can be expensive.
	// TODO: How to know which repo is the one for the PT project? Should they be named the same?
	client := github.NewClient(nil)
	w.WriteHeader(http.StatusOK)
}
