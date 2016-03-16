package server

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/BaristaVentures/errand-boy/config"
	"github.com/BaristaVentures/errand-boy/routers/repos"
	"github.com/BaristaVentures/errand-boy/routers/tracker"
	// Importing it as a blank package causes its init method to be called.
	"github.com/BaristaVentures/errand-boy/services/repotracker"
	trackerservice "github.com/BaristaVentures/errand-boy/services/tracker"
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
	trackerSubRouter := hooksSubRouter.PathPrefix("/pt").Subrouter()

	// Add the repos.NormalizePRPayload middleware.
	baseRoute.Handler(repos.NormalizePRPayload(hooksSubRouter))
	// Add the repos routes.
	repos.Route(reposSubRouter)
	//Add the tracker routes.
	tracker.Route(trackerSubRouter)

	apiTokenEnvVar := config.Current().TrackerAPIToken
	if len(apiTokenEnvVar) == 0 {
		panic(errors.New("Pivotal Tracker API Token Environment Variable name not set in config file."))
	}
	trackerAPIToken := os.Getenv(apiTokenEnvVar)
	if len(trackerAPIToken) == 0 {
		panic(errors.New("No Pivotal Tracker API Token found in env var " + apiTokenEnvVar))
	}
	service := trackerservice.New(trackerAPIToken)
	repotracker.SetTrackerService(service)

	http.ListenAndServe(":"+strconv.Itoa(s.Port), r)
}
