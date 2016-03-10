package tracker

import (
	"encoding/json"
	"net/http"

	"github.com/BaristaVentures/errand-boy/routers"
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
	w.WriteHeader(http.StatusOK)
}
