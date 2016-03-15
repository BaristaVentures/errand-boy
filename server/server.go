package server

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	// Importing it as a blank package causes its init method to be called.
	"github.com/BaristaVentures/errand-boy/services/repotracker"
	"github.com/BaristaVentures/errand-boy/services/tracker"
	"github.com/gorilla/mux"
)

// Server represents the server's config.
type Server struct {
	Port int
}

// BootUp starts the server.
func (s *Server) BootUp() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	baseRoute := r.PathPrefix("/")

	hooksSubRoute := r.PathPrefix("/hooks")
	hooksSubRouter := hooksSubRoute.Subrouter()

	reposSubRouter := hooksSubRouter.PathPrefix("/repos").Subrouter()

	baseRoute.Handler(repos.NormalizePRPayload(hooksSubRouter))

	repos.Route(reposSubRouter)

	apiToken := config.Current().TrackerAPIToken
	if len(apiToken) == 0 {
		panic(errors.New("Pivotal Tracker API Token not set in config file."))
	}
	service := tracker.NewService(os.Getenv(apiToken))
	repotracker.SetTrackerService(service)

	http.ListenAndServe(":"+strconv.Itoa(s.Port), r)
}
