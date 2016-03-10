package repos

import (
	"errors"

	"github.com/BaristaVentures/errand-boy/routers"
	"github.com/BaristaVentures/errand-boy/utils"
	"github.com/gorilla/mux"
)

var events = []string{"pr"}
var eventsSubs = make(map[string]*utils.Observers)

func init() {
	for _, ev := range events {
		observers := make(utils.Observers, 0)
		eventsSubs[ev] = &observers
	}
}

// AddObserver adds an observer to the list.
func AddObserver(event string, observer utils.Observer) error {
	_, ok := eventsSubs[event]
	if !ok {
		return errors.New("No such event: " + event)
	}
	eventsSubs[event].AddObserver(observer)
	return nil
}

// Route returns a configured *mux.Router.
func Route(router *mux.Router) *mux.Router {
	routes := &routers.Routes{
		&routers.Route{
			Path:    "/pr",
			Method:  "POST",
			Handler: pullRequestHandler,
		},
	}

	for _, r := range *routes {
		router.Methods(r.Method).Path(r.Path).Handler(r.Handler)
	}
	return router
}
